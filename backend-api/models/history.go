package models

import "time"

type HistoryLog struct {
	ID          string `json:"id" gorm:"primaryKey"`
	UserId      string `json:"user_id" gorm:"not null"`
	User        User
	Type        string    `json:"type" gorm:"type:enum('general', 'workspace')"`
	UserRole    string    `json:"user_role" gorm:"type:enum('admin', 'user')"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserHistoryLog struct {
	HistoryId string `json:"history_id" gorm:"primaryKey"`
	History   HistoryLog
	UserId    string `json:"user_id" gorm:"primaryKey"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WorkspaceHistoryLog struct {
	HistoryId   string `json:"history_id" gorm:"primaryKey"`
	History     HistoryLog
	WorkspaceId string `json:"workspace_id" gorm:"primaryKey"`
	Workspace   Workspace
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
