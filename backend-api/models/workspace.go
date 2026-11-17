package models

import "time"

type Workspace struct {
	Id            string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name          string `json:"name" gorm:"not null;type:varchar(255)"`
	Description   string `json:"description" gorm:"type:text"`
	UserCreatedID string `json:"user_id" gorm:"not null;type:varchar(36)"`
	UserCreated   User
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type WorkspaceMember struct {
	WorkspaceID string `json:"workspace_id" gorm:"primaryKey;type:varchar(36)"`
	Workspace   Workspace
	UserID      string `json:"user_id" gorm:"primaryKey;type:varchar(36)"`
	User        User
	Role        string    `json:"role" gorm:"type:enum('admin', 'member');default:'member';not null"`
	JoinedAt    time.Time `json:"joined_at" gorm:"autoCreateTime"`
}

type WorkspaceCardStatus struct {
	Id          string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Order       int    `json:"order" gorm:"not null"`
	WorkspaceID string `json:"workspace_id" gorm:"not null;type:varchar(36)"`
	Workspace   Workspace
	Name        string    `json:"name" gorm:"not null;type:varchar(255)"`
	Color       string    `json:"color" gorm:"not null;type:varchar(255)"`
	IsArchived  bool      `json:"is_archived" gorm:"default:false"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WorkspaceCard struct {
	Id            string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Order         int    `json:"order" gorm:"not null"`
	WorkspaceID   string `json:"workspace_id" gorm:"not null;type:varchar(36)"`
	Workspace     Workspace
	StatusID      string `json:"status_id" gorm:"not null;type:varchar(36)"`
	Status        WorkspaceCardStatus
	UserCreatedID string `json:"user_id" gorm:"not null;type:varchar(36)"`
	UserCreated   User
	Title         string    `json:"title" gorm:"not null;type:varchar(255)"`
	Description   string    `json:"description;type:text"`
	IsArchived    bool      `json:"is_archived" gorm:"default:false"`
	IsCompleted   bool      `json:"is_completed" gorm:"default:false"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type WorkspaceCardMember struct {
	CardID   string `json:"card_id" gorm:"primaryKey;type:varchar(36)"`
	Card     WorkspaceCard
	UserID   string `json:"user_id" gorm:"primaryKey;type:varchar(36)"`
	User     User
	JoinedAt time.Time `json:"joined_at" gorm:"autoCreateTime"`
}

type WorkspaceCardActivity struct {
	Id        string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CardID    string `json:"card_id" gorm:"not null;type:varchar(36)"`
	Card      WorkspaceCard
	UserID    string `json:"user_id" gorm:"not null;type:varchar(36)"`
	User      User
	Type      string    `json:"type" gorm:"type:enum('general', 'comment')"`
	Message   string    `json:"message" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
