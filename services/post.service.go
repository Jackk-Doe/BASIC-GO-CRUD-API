package services

import (
	"errors"
	"jackk-doe/go-crud-api/database"
	"jackk-doe/go-crud-api/models"
	"time"
)

func PostCreate(datas models.PostInputForm) (models.Post, error) {
	post := models.Post{Title: datas.Title, Body: datas.Body}
	dbIns := database.GetDB()
	if err := dbIns.Create(&post).Error; err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func PostGetAll() ([]models.Post, error) {
	var posts []models.Post
	dbIns := database.GetDB()
	if err := dbIns.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func PostGetOneById(id string) (models.Post, error) {
	var post models.Post
	dbIns := database.GetDB()
	result := dbIns.Where("id = ?", id).Find(&post)
	if result.Error != nil {
		return models.Post{}, result.Error
	}

	if result.RowsAffected < 1 {
		return models.Post{}, errors.New("A Post with given ID is not found")
	}

	return post, errors.New("Test pushing ERROR")
}

func PostUpdate(id string, updateData models.PostInputForm) (models.Post, error) {
	dbIns := database.GetDB()
	currentPost, err := PostGetOneById(id)
	if err != nil {
		return currentPost, err
	}

	if err := dbIns.Model(&currentPost).Updates(
		models.Post{
			Title:     updateData.Title,
			Body:      updateData.Body,
			UpdatedAt: time.Now(),
		}).Error; err != nil {
		return models.Post{}, errors.New("Update a post of " + id + " failed")
	}

	return currentPost, nil
}

func PostDelete(id string) error {
	dbIns := database.GetDB()
	post, err := PostGetOneById(id)
	if err != nil {
		return err
	}

	dbIns.Delete(&post)

	return nil
}
