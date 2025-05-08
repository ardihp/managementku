package models

import "time"

type Workspace struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceMember struct {
	WorkspaceId int `json:"workspaceId"`
	Workspace   Workspace
	UserId      int `json:"user_id"`
	User        User
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceTag struct {
	WorkspaceId int `json:"workspaceId"`
	Workspace   Workspace
	Tag         string `json:"tag"`
}
