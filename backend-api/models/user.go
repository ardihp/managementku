package models

import "time"

type User struct {
	Id        string     `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string     `json:"name" gorm:"type:varchar(255)"`
	Username  string     `json:"username" gorm:"unique;not null;type:varchar(50)"`
	Email     string     `json:"email" gorm:"unique;not null;type:varchar(50)"`
	Password  string     `json:"password" gorm:"type:varchar(50)"`
	IsDeleted bool       `json:"is_deleted" gorm:"default:false"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
