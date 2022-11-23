package services

import (
	"errors"
	"jackk-doe/go-crud-api/database"
	"jackk-doe/go-crud-api/models"
	"time"

	"gorm.io/gorm"
)

func PostCreate(datas models.PostCreate) (models.Post, error) {
	post := models.Post{Title: datas.Title, Body: datas.Body}
	dbIns := database.GetDB()
	result := dbIns.Create(&post)

	if result.Error != nil {
		return models.Post{}, errors.New(result.Error.Error())
	}

	return post, nil
}

func PostGetAll() ([]models.Post, error) {
	var posts []models.Post
	dbIns := database.GetDB()
	result := dbIns.Find(&posts)

	if result.Error != nil {
		return nil, errors.New(result.Error.Error())
	}

	return posts, nil
}

func PostGetOneById(id string) (models.Post, error) {
	var post models.Post
	dbIns := database.GetDB()

	if result := dbIns.Where("id = ?", id).First(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return post, errors.New("Given Id is not found")
		}
		return post, errors.New(result.Error.Error())
	}

	return post, nil
}

func PostUpdate(id string, updateData models.PostCreate) (models.Post, error) {
	dbIns := database.GetDB()
	currentPost, err := PostGetOneById(id)
	if err != nil {
		return currentPost, errors.New(err.Error())
	}

	if result := dbIns.Model(&currentPost).Updates(models.Post{Title: updateData.Title, Body: updateData.Body, UpdatedAt: time.Now()}); result.Error != nil {
		return currentPost, errors.New("Update a post of " + id + " failed")
	}

	return currentPost, nil
}

func PostDelete(id string) error {
	dbIns := database.GetDB()
	post, err := PostGetOneById(id)
	if err != nil {
		return errors.New(err.Error())
	}

	dbIns.Delete(&post)

	return nil
}
