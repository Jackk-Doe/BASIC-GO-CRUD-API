package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// [Post] for interact with database,
// for safety DO NOT return this struct to client
type Post struct {
	ID        string `gorm:"primarKey"`
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
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// This function is run before [Post] is created into database
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	idString := uuid.New().String()
	p.ID = idString
	return
}

// Convert from [Post] struct to [PostDTO] struct
func (p *Post) ToDto() PostDTO {
	return PostDTO{p.ID, p.Title, p.Body, p.CreatedAt, p.UpdatedAt}
}
