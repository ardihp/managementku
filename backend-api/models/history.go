package models

import "time"

type HistoryLog struct {
	ID          string `json:"id" gorm:"primaryKey;type:varchar(36)"`
	UserID      string `json:"user_id" gorm:"not null;type:varchar(36)"`
	User        User
	Type        string    `json:"type" gorm:"type:enum('general', 'workspace')"`
	UserRole    string    `json:"user_role" gorm:"type:enum('admin', 'user')"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserHistoryLog struct {
	HistoryID string `json:"history_id" gorm:"primaryKey;type:varchar(36)"`
	History   HistoryLog
	UserID    string `json:"user_id" gorm:"primaryKey;type:varchar(36)"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceHistoryLog struct {
	HistoryID   string `json:"history_id" gorm:"primaryKey;type:varchar(36)"`
	History     HistoryLog
	WorkspaceID string `json:"workspace_id" gorm:"primaryKey;type:varchar(36)"`
	Workspace   Workspace
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
