package controllers

import (
	"net/http"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/services"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var post models.PostInputForm

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPost, err := services.PostCreate(post)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to PostDTO
	postDto := createdPost.ToDto()

	c.JSON(http.StatusCreated, gin.H{
		"post": postDto,
	})
}

func PostGetAll(c *gin.Context) {
	posts, err := services.PostGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to list of Post DTO
	postsDto := make([]models.PostDTO, len(posts))
	for i, post := range posts {
		postsDto[i] = post.ToDto()
	}

	c.JSON(http.StatusOK, gin.H{"posts": postsDto})
}

func PostGetOneById(c *gin.Context) {
	id := c.Param("id")
	post, err := services.PostGetOneById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Convert to PostDTO
	postDto := post.ToDto()
	c.JSON(http.StatusOK, gin.H{"post": postDto})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var updateData models.PostInputForm
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPost, err := services.PostUpdate(id, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert to PostDTO
	postDto := updatedPost.ToDto()

	c.JSON(http.StatusCreated, gin.H{
		"post": postDto,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	if err := services.PostDelete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted success"})
}
