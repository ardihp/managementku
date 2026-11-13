package models

import "time"

type User struct {
	Id        uint       `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	Username  string     `json:"username" gorm:"unique;not null"`
	Email     string     `json:"email" gorm:"unique;not null"`
	Password  string     `json:"password"`
	IsDeleted bool       `json:"is_deleted" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
