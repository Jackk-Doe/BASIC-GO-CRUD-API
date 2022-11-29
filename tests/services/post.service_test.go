package services_test

import (
	"log"
	"testing"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/services"

	"github.com/stretchr/testify/assert"
)

func Test_Post_Services_Functions(t *testing.T) {
	log.Println()
	log.Println("--> START : Post Services unit tests...")
	log.Println()

	/// Mock Post data Input Form
	mockPost := models.PostInputForm{Title: "Test title", Body: "This is a body of Test Post"}

	// To track if the mock data is created successfully
	isPostCreateSuccess := false
	var createdPostId string

	/// Mock User model data
	mockUser := models.User{ID: "This-is-mock-user-ID", Email: "mock-user@email.com", Name: "mock-username", Password: "mockUserPassword"}

	t.Run("TEST_CREATE_POST", func(t *testing.T) {
		createdPost, err := services.PostCreate(mockPost, mockUser)

		assert.Equal(t, err, nil, "Create Post : Must not return Error")
		assert.Equal(t, mockPost.Title, createdPost.Title, "Create Post : created Title must be the same with input Title")
		assert.Equal(t, mockPost.Body, createdPost.Body, "Create Post : created Body must be the same with input Body")

		// If created success, re-assign
		if err == nil {
			isPostCreateSuccess = true
			createdPostId = createdPost.ID
		}
	})

	t.Run("TEST_GET_ALL_POSTS", func(t *testing.T) {
		posts, err := services.PostGetAll()

		assert.Equal(t, err, nil, "Get All Posts : Must not return Error")
		assert.NotEqual(t, posts, nil, "Get All Posts : Return Posts datas must not be NULL")
	})

	/*
		NOTE : Only runs the below Unit tests, if the Mock Post is created SUCCESSFULLY
	*/
	if isPostCreateSuccess {

		// Later : assign Post from PostGetOneById() to this var
		var post models.Post

		t.Run("TEST_GET_POST_BY_ID", func(t *testing.T) {
			var getErr error
			post, getErr = services.PostGetOneById(createdPostId)

			assert.Equal(t, getErr, nil, "Get Post by ID : Must not return Error")
			assert.Equal(t, mockPost.Title, post.Title, "Get Post by ID : returned Title must be the same with input Title")
			assert.Equal(t, mockPost.Body, post.Body, "Get Post by ID : returned Body must be the same with input Body")
		})

		t.Run("TEST_UPDATE_POST_BY_ID", func(t *testing.T) {
			updateInput := models.PostInputForm{Title: "UPDATE title", Body: "This is a body of UPDATE Post"}

			updated, err := services.PostUpdate(createdPostId, post, updateInput)

			assert.Equal(t, err, nil, "Update Post by ID : Must not return Error")
			assert.Equal(t, updateInput.Title, updated.Title, "Update Post by ID : updated Title must be the same with new input Title")
			assert.Equal(t, updateInput.Body, updated.Body, "Update Post by ID : updated Body must be the same with new input Body")
		})

		t.Run("TEST_DELETE_POST_BY_ID", func(t *testing.T) {
			err := services.PostDelete(post)

			assert.Equal(t, err, nil, "Delete Post by ID : Must not return Error")
		})
	}

	log.Println()
	log.Println("--> END : Post Services unit tests...")
	log.Println()
}
