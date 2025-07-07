package models

import "time"

type HistoryLog struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id" gorm:"not null"`
	User        User
	Type        string    `json:"type" gorm:"type:enum('general', 'workspace')"`
	Action      string    `json:"action"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserHistoryLog struct {
	HistoryId uint `json:"history_id" gorm:"not null"`
	History   HistoryLog
	UserId    uint `json:"user_id" gorm:"not null"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
