package model

import "time"

// User struct
type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(256);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
