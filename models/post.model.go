package models

import (
	"time"
)

/// Old Version
// type Post struct {
// 	gorm.Model
// 	Title string
// 	Body  string
// }

type Post struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}

// [PostCreate] when a Post is created from POST & PUT method
type PostCreate struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TODO : use [PostRead] instead of [Post] to return posts data
// [PostRead] when a server recieved query GET method
// type PostRead struct {
// 	ID    string `json:"id" gorm:"primary_key"`
// 	Title string `json:"title"`
// 	Body  string `json:"body"`
// 	CreatedAt time.Time `json:"createAt"`
// 	UpdatedAt time.Time `json:"updatedAt"`
// }
