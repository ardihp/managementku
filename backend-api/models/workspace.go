package models

import "time"

type Workspace struct {
	Id              uint   `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	WorkspaceTypeId uint   `json:"type_id"`
	WorkspaceType   WorkspaceType
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type WorkspaceRole struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceMember struct {
	WorkspaceId uint `json:"workspace_id"`
	Workspace   Workspace
	UserId      uint `json:"user_id"`
	User        User
	RoleId      uint `json:"role_id"`
	Role        WorkspaceRole
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceTag struct {
	WorkspaceId uint `json:"workspace_id"`
	Workspace   Workspace
	Tag         string    `json:"tag"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceType struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceStatus struct {
	WorkspaceId uint `json:"workspace_id"`
	Workspace   Workspace
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
