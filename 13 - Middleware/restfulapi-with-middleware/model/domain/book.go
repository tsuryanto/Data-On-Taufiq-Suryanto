package domain

import "gorm.io/gorm"

type Book struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
	gorm.Model
}
