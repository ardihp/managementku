package controllers

import (
	"encoding/json"
	"net/http"
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
	Text string `json:"text"`
	Type string `json:"type"`
}

type Message struct {
	Sender   string         `json:"sender,omitempty"`
	Receiver string         `json:"receiver,omitempty"`
	Content  MessageContent `json:"content,omitempty"`
}

var manager = &ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		if conn != ignore {
			conn.Send <- message
		}
	}
}

// StartManager runs the client manager loop.
func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: MessageContent{Text: "A new socket has connected.", Type: "initialize"}})
			manager.send(jsonMessage, conn)

		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: MessageContent{Text: "A socket has disconnected.", Type: "initialize"}})
				manager.send(jsonMessage, conn)
			}

		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

func (c *Client) read() {
	defer func() {
		manager.Unregister <- c
		_ = c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()

		var parsedMessage MessageContent

		json.Unmarshal([]byte(string(message)), &parsedMessage)

		if err != nil {
			manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}

		time.Sleep(600 * time.Millisecond)
		jsonMessage, _ := json.Marshal(&Message{Sender: c.Id, Content: MessageContent{Text: parsedMessage.Text, Type: parsedMessage.Type}})
		manager.Broadcast <- jsonMessage
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
	go manager.start()

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	client := &Client{
		Id:     c.GetString("user_id"),
		Socket: ws,
		Send:   make(chan []byte),
	}

	manager.Register <- client

	go client.read()
	go client.write()
}
