package dto

import "time"

/* REQUEST */
type UserRequestBody struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserRequestAuth struct {
	Identifier string `validate:"required" json:"identifier"`
	Password   string `validate:"required" json:"password"`
}

/* RESPONSE */
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
