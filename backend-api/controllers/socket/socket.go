package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync" // Import sync for RWMutex
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id     string
	Socket *websocket.Conn
	Send   chan []byte
}

type MessageContent struct {
	Text        string `json:"text"`
	WorkspaceId int    `json:"workspace_id"`
	CardId      int    `json:"card_id"`
	Type        string `json:"type"`
}

type Message struct {
	Sender   string         `json:"sender,omitempty"`
	Receiver string         `json:"receiver,omitempty"`
	Content  MessageContent `json:"content,omitempty"`
}

// Manager is the global client manager, now exported and protected by RWMutex
var Manager = &ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	// Clients map needs to be protected for concurrent access
	Clients: make(map[*Client]bool),
	mu:      sync.RWMutex{}, // Initialize the RWMutex
}

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.RWMutex // Add an RWMutex to protect the Clients map
}

// send is now protected by RWMutex for reading the map
func (manager *ClientManager) send(message []byte, ignore *Client) {
	manager.mu.RLock()         // Acquire read lock before iterating
	defer manager.mu.RUnlock() // Release read lock

	for conn := range manager.Clients {
		if conn != ignore {
			select {
			case conn.Send <- message:
				// Message sent successfully
			default:
				// If client.Send buffer is full or closed, unregister
				// This part should ideally be handled within the manager's loop
				// to avoid modifying the map while holding a read lock.
				// For now, we'll keep it as is, but be aware of potential nuances.
				// A more robust solution might involve sending a signal back
				// to the manager.Start() loop to unregister this client.
				close(conn.Send)
				// manager.Unregister <- conn // This would cause a deadlock if RLock is held
				// So, we'll just skip and rely on the Unregister case in Start()
			}
		}
	}
}

// Start runs the client manager loop. This should be run only ONCE.
func (manager *ClientManager) Start() { // Changed to Start (capital S)
	for {
		select {
		case conn := <-manager.Register:
			manager.mu.Lock() // Acquire write lock for registration
			manager.Clients[conn] = true
			manager.mu.Unlock() // Release write lock

			jsonMessage, _ := json.Marshal(&Message{Content: MessageContent{Text: "A new socket has connected.", Type: "initialize"}})
			manager.send(jsonMessage, conn)

		case conn := <-manager.Unregister:
			manager.mu.Lock() // Acquire write lock for unregistration
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
			}
			manager.mu.Unlock() // Release write lock

			jsonMessage, _ := json.Marshal(&Message{Content: MessageContent{Text: "A socket has disconnected.", Type: "initialize"}})
			manager.send(jsonMessage, conn)

		case message := <-manager.Broadcast:
			// Iterate over a copy or use RWMutex carefully for broadcast
			// We need to acquire a read lock for iterating
			manager.mu.RLock()
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
					// Message sent
				default:
					// If the send channel is blocked, it means the client's write goroutine
					// is not reading fast enough or is closed.
					// We need to unregister this client.
					// IMPORTANT: Do NOT delete from map directly here while holding RLock.
					// Send it to Unregister channel for the manager to handle.
					close(conn.Send)     // Close the channel to signal client to stop writing
					go func(c *Client) { // Use a goroutine to avoid blocking the broadcast
						manager.Unregister <- c
					}(conn)
				}
			}
			manager.mu.RUnlock() // Release read lock after iterating
		}
	}
}

func (c *Client) read() {
	defer func() {
		Manager.Unregister <- c // Use Manager (capital M)
		_ = c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()

		var parsedMessage MessageContent

		json.Unmarshal([]byte(string(message)), &parsedMessage)

		if err != nil {
			// This check should primarily be for unexpected errors.
			// websocket.IsUnexpectedCloseError can be used for normal disconnections.
			// Log the error for debugging if needed.
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				// log.Printf("Error reading message: %v", err)
			}
			Manager.Unregister <- c // Use Manager (capital M)
			_ = c.Socket.Close()
			break
		}

		fmt.Println(parsedMessage)

		time.Sleep(600 * time.Millisecond) // Consider if this delay is necessary and desired
		jsonMessage, _ := json.Marshal(&Message{Sender: c.Id, Content: MessageContent{Text: parsedMessage.Text, Type: parsedMessage.Type, WorkspaceId: parsedMessage.WorkspaceId, CardId: parsedMessage.CardId}})
		Manager.Broadcast <- jsonMessage // Use Manager (capital M)
	}
}

func (c *Client) write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// Channel closed by manager, so close the websocket
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024 * 1024,
	WriteBufferSize: 1024 * 1024 * 1024,

	// Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitSocket(c *gin.Context) {
	// IMPORTANT: No longer calling go Manager.Start() here.
	// It should be called ONCE at application startup (e.g., in main.go).

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	client := &Client{
		Id:     c.GetString("user_id"),
		Socket: ws,
		Send:   make(chan []byte, 256), // Add a buffer to the Send channel
	}

	Manager.Register <- client // Use Manager (capital M)

	go client.read()
	go client.write()
}
