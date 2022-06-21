package domain

import (
	"gorm.io/gorm"
)

type User struct {
	// ID        uint      `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt time.Time `json:"deleted_at"`
	gorm.Model
}
