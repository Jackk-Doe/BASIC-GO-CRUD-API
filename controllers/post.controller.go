package controllers

import (
	"errors"
	"net/http"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/services"

	"github.com/gin-gonic/gin"
)

// A Helper function to get User, via [user_id] from gin Context
func getUserFromGinContext(c *gin.Context) (models.User, error) {
	userId, isFoundID := c.Get("user_id")
	if isFoundID == false {
		return models.User{}, errors.New("Field [user_id] can not be found inside the request")
	}
	// NOTE : Convert form any -> string
	convertedUserId := userId.(string)

	// Get User via ID, also check User existing
	user, isFoundUser, findErr := services.UserGetViaID(convertedUserId)
	if findErr != nil {
		return models.User{}, findErr
	}
	if isFoundUser == false {
		return models.User{}, errors.New("User with the given ID is not found")
	}
	return user, nil
}

func PostCreate(c *gin.Context) {
	var post models.PostInputForm

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get User (Author) from gin Context
	author, getAuthorErr := getUserFromGinContext(c)
	if getAuthorErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getAuthorErr.Error()})
		return
	}

	// Check if the input Title is already existed
	if titleExisted, err := services.TitleExisted(post.Title); titleExisted == true {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Title is already existed"})
		return
	}

	createdPost, err := services.PostCreate(post, author)

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

	// Check if a Post of the given id existed
	currentPost, err := services.PostGetOneById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the input Title is already existed
	if titleExisted, err := services.TitleExisted(updateData.Title); titleExisted == true {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "The Title is already existed"})
		return
	}

	updatedPost, err := services.PostUpdate(id, currentPost, updateData)
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
