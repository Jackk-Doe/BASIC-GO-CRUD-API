package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}

// [PostCreate] when a Post is created from POST method
type PostCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
