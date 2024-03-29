package dto

import "time"

/* REQUEST */
type UserRequestBody struct {
	Name     string `validate:"required" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UpdateUserRequestBody struct {
	Name     *string `json:"name"`
	Email    *string `validate:"omitempty,email" json:"email"`
	Password *string `json:"password"`
}

/* RESPONSE */
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
