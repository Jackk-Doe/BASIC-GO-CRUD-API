package services

import (
	"errors"
	"time"

	"github.com/Jackk-Doe/basic-go-crud-api/database"
	"github.com/Jackk-Doe/basic-go-crud-api/models"
)

func PostCreate(datas models.PostInputForm, author models.User) (models.Post, error) {
	post := models.Post{Title: datas.Title, Body: datas.Body, Author: author}
	dbIns := database.GetDB()
	if err := dbIns.Create(&post).Error; err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func PostGetAll() ([]models.Post, error) {
	var posts []models.Post
	dbIns := database.GetDB()
	if err := dbIns.Model(&models.Post{}).Preload("Author").Find(&posts).Error; err != nil {
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

	return post, nil
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

func TitleExisted(title string) (bool, error) {
	dbIns := database.GetDB()
	result := dbIns.Where("title = ?", title).Find(&models.Post{})
	if result.Error != nil {
		return true, result.Error //! Error : Finding
	}
	if result.RowsAffected > 0 {
		return true, nil //Already EXISTED
	}
	return false, nil //Not found
}
