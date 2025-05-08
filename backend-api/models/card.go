package models

import "time"

type Card struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	WorkspaceId int  `json:"workspaceId"`
	Workspace   Workspace
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CardMember struct {
	CardId    int `json:"cardId"`
	Card      Card
	UserId    int `json:"userId"`
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
