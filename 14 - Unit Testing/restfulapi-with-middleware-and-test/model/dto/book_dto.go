package dto

import "time"

/* REQUEST */
type BookRequestBody struct {
	Name      string `validate:"required" json:"name"`
	Publisher string `validate:"required" json:"publisher"`
	Author    string `validate:"required" json:"author"`
}

/* RESPONSE */
type BookResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Publisher string    `json:"publisher"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
