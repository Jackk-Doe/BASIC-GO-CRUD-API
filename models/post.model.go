package models

import (
	"time"
)

// [Post] for interact with database,
// for safety DO NOT return this struct to client
type Post struct {
	ID        uint   `gorm:"primarKey"`
	Title     string `gorm:"unique"`
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// [PostInputForm] when a Post is created from POST & PUT method
type PostInputForm struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// [PostDTO] for datas transfer object (sending to clients)
type PostDTO struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Convert from [Post] struct to [PostDTO] struct
func (p *Post) ToDto() PostDTO {
	return PostDTO{p.ID, p.Title, p.Body, p.CreatedAt, p.UpdatedAt}
}
