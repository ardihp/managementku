package models

import "time"

type Workspace struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	Name          string `json:"name" gorm:"not null"`
	Description   string `json:"description"`
	UserCreatedId uint   `json:"user_id" gorm:"not null"`
	UserCreated   User
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type WorkspaceMember struct {
	WorkspaceId uint `json:"workspace_id" gorm:"primaryKey"`
	Workspace   Workspace
	UserId      uint `json:"user_id" gorm:"primaryKey"`
	User        User
	Role        string    `json:"role" gorm:"default:'user'"`
	JoinedAt    time.Time `json:"joined_at" gorm:"autoCreateTime"`
}

type WorkspaceCardStatus struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	Order       int  `json:"order" gorm:"not null"`
	WorkspaceId uint `json:"workspace_id" gorm:"not null"`
	Workspace   Workspace
	Name        string    `json:"name"`
	Color       string    `json:"color"`
	IsArchived  bool      `json:"is_archived" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceCard struct {
	Id                    uint `json:"id" gorm:"primaryKey"`
	Order                 int  `json:"order" gorm:"not null"`
	WorkspaceId           uint `json:"workspace_id" gorm:"not null"`
	Workspace             Workspace
	WorkspaceCardStatusId uint `json:"workspace_card_status_id" gorm:"not null"`
	WorkspaceCardStatus   WorkspaceCardStatus
	UserCreatedId         uint `json:"user_id" gorm:"not null"`
	UserCreated           User
	Title                 string    `json:"title" gorm:"not null"`
	Description           string    `json:"description"`
	IsArchived            bool      `json:"is_archived" gorm:"default:false"`
	IsCompleted           bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type WorkspaceCardMember struct {
	WorkspaceCardId uint `json:"workspace_card_id" gorm:"primaryKey"`
	WorkspaceCard   WorkspaceCard
	UserId          uint `json:"user_id" gorm:"primaryKey"`
	User            User
	JoinedAt        time.Time `json:"joined_at"`
}

type WorkspaceCardActivity struct {
	Id              uint `json:"id" gorm:"primaryKey"`
	WorkspaceCardId uint `json:"workspace_card_id" gorm:"not null"`
	WorkspaceCard   WorkspaceCard
	UserId          uint `json:"user_id" gorm:"not null"`
	User            User
	Type            string    `json:"type" gorm:"type:enum('general', 'comment')"`
	Message         string    `json:"message"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
