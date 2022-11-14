package services

import (
	"errors"
	"jackk-doe/go-crud-api/initializers"
	"jackk-doe/go-crud-api/models"
)

func PostCreate(datas models.PostCreate) (models.Post, error) {
	post := models.Post{Title: datas.Title, Body: datas.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		return models.Post{}, errors.New(result.Error.Error())
	}

	return post, nil
}

func PostGetAll() ([]models.Post, error) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}

	return posts, nil
}
