package structs

import "kanban-board/backend-api/models"

type WorkspaceResponse struct {
	Id          uint          `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Members     []models.User `json:"members"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}
type WorkspaceCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Members     []int  `json:"members"`
}

type WorkspaceUpdateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Category    string `json:"category,omitempty"`
	Members     []int  `json:"members"`
}

type WorkspaceAssignMember struct {
	WorkspaceId int   `json:"workspace_id" binding:"required"`
	UserId      []int `json:"user_id" binding:"required"`
}
