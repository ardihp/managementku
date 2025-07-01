package models

import "time"

type Workspace struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	Order       int    `json:"order" gorm:"not null"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TypeId      uint   `json:"type_id"`
	Type        WorkspaceType
	StatusId    uint `json:"status_id"`
	Status      WorkspaceStatus
	IsArchived  bool      `json:"is_archived" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceMember struct {
	WorkspaceId uint `json:"workspace_id"`
	Workspace   Workspace
	UserId      uint `json:"user_id"`
	User        User
	Role        string
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
	Id         uint      `json:"id" gorm:"primaryKey"`
	Order      int       `json:"order" gorm:"not null"`
	Name       string    `json:"name"`
	Color      string    `json:"color"`
	IsArchived bool      `json:"is_archived" gorm:"default:false"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type WorkspaceCardStatus struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	Order       int  `json:"order" gorm:"not null"`
	WorkspaceId uint `json:"workspace_id"`
	Workspace   Workspace
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	IsArchived  bool      `json:"is_archived" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceCard struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	Order       int  `json:"order" gorm:"not null"`
	WorkspaceId uint `json:"workspace_id"`
	Workspace   Workspace
	StatusId    uint `json:"status_id"`
	Status      WorkspaceCardStatus
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsArchived  bool      `json:"is_archived" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceCardMember struct {
	CardId    uint `json:"card_id"`
	Card      WorkspaceCard
	UserId    uint `json:"user_id"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
