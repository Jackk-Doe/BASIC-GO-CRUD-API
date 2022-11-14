package controllers

import (
	"jackk-doe/go-crud-api/models"
	"jackk-doe/go-crud-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var post models.PostCreate

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Post.Title must not be empty
	if post.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post' Title can not be empty"})
		return
	}

	createdPost, err := services.PostCreate(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": createdPost,
	})
}
