package structs

import "time"

type MemberDetail struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type WorkspaceMember struct {
	WorkspaceId string `json:"workspace_id"`
	UserId      string `json:"user_id"`
	Role        string `json:"role"`
}

type WorkspaceDetailResponse struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Members     []MemberDetail `json:"members"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type WorkspaceCardStatus struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	IsArchived bool   `json:"is_archived"`
	Order      int    `json:"order"`
}

type WorkspaceCardActivityResponse struct {
	Id          string `json:"id"`
	UserCreated string `json:"user_created"`
	Type        string `json:"type"`
	Message     string `json:"message"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type WorkspaceCardResponse struct {
	Id           string         `json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Status       string         `json:"status"`
	UserCreated  string         `json:"user_created"`
	Members      []MemberDetail `json:"members"`
	CommentCount int            `json:"comment_count"`
	IsArchived   bool           `json:"is_archived"`
	IsCompleted  bool           `json:"is_completed"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
}

type CountWorkspaceCard struct {
	Completed   string `json:"completed"`
	Incompleted string `json:"incomplete"`
}

type WorkspaceResponse struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Members     []MemberDetail `json:"members"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type CreateWorkspaceResponse struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type WorkspaceUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type WorkspaceAssignMember struct {
	UserId []string `json:"user_id" binding:"required"`
}

type WorkspaceCardAssignMember struct {
	CardId string   `json:"card_id" binding:"required"`
	UserId []string `json:"user_id" binding:"required"`
}

type WorkspaceWatchCardRequest struct {
	UserId string `json:"user_id" binding:"required"`
}

type WorkspaceCreateRoleRequest struct {
	Name string `json:"name" binding:"required"`
}

type WorkspaceCreateTypeRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateTypeRequest struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateTypeOrderRequest struct {
	OrderStatus []string `json:"order_status" binding:"required"`
}

type WorkspaceCardByStatusResponse struct {
	Id         string                  `json:"id"`
	Name       string                  `json:"name"`
	Color      string                  `json:"color"`
	IsArchived bool                    `json:"is_archived"`
	Cards      []WorkspaceCardResponse `json:"cards"`
}

type WorkspaceCreateCardRequest struct {
	StatusId    string `json:"status_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type WorkspaceUpdateCardRequest struct {
	StatusId    string `json:"status_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type WorkspaceCreateStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceCreateCardStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateStatusRequest struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceUpdateCardStatusRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

type WorkspaceOrder struct {
	StatusId    string   `json:"status_id" binding:"required"`
	WorkspaceId []string `json:"workspace_id" binding:"required"`
}

type WorkspaceUpdateOrderRequest struct {
	WorkspaceId string           `json:"workspace_id" binding:"required"`
	OrderStatus []WorkspaceOrder `json:"order_status" binding:"required"`
}

type WorkspaceCardUpdateOrderStatusRequest struct {
	OrderStatus []string `json:"order_status" binding:"required"`
}

type WorkspaceCardOrder struct {
	StatusId string   `json:"status_id" binding:"required"`
	CardId   []string `json:"card_id" binding:"required"`
}

type WorkspaceCardUpdateOrderRequest struct {
	CardId      string               `json:"card_id" binding:"required"`
	OrderStatus []WorkspaceCardOrder `json:"order_status" binding:"required"`
}

type WorkspaceCardCreateActivityRequst struct {
	Message string `json:"message"`
}
