package services_test

import (
	"strconv"
	"testing"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/services"

	"github.com/stretchr/testify/assert"
)

func Test_Post_Services_Functions(t *testing.T) {

	// Test with this mock data
	mockInput := models.PostInputForm{Title: "Test title", Body: "This is a body of Test Post"}

	// To track if the mock data is created successfully
	isCreateSucc := false
	postId := uint(1)

	t.Run("TEST_CREATE_POST", func(t *testing.T) {
		createdPost, err := services.PostCreate(mockInput)

		assert.Equal(t, err, nil, "Create Post : Must not return Error")
		assert.Equal(t, mockInput.Title, createdPost.Title, "Create Post : created Title must be the same with input Title")
		assert.Equal(t, mockInput.Body, createdPost.Body, "Create Post : created Body must be the same with input Body")

		// If created success, re-assign
		if err == nil {
			isCreateSucc = true
			postId = createdPost.ID
		}
	})

	t.Run("TEST_GET_ALL_POSTS", func(t *testing.T) {
		posts, err := services.PostGetAll()

		assert.Equal(t, err, nil, "Get All Posts : Must not return Error")
		assert.NotEqual(t, posts, nil, "Get All Posts : Return Posts datas must not be NULL")
	})

	/*
		NOTE : Only runs the below Unit tests, if the mock datas is created SUCCESSFULLY
	*/
	if isCreateSucc {

		t.Run("TEST_GET_POST_BY_ID", func(t *testing.T) {
			post, err := services.PostGetOneById(strconv.Itoa(int(postId)))

			assert.Equal(t, err, nil, "Get Post by ID : Must not return Error")
			assert.Equal(t, mockInput.Title, post.Title, "Get Post by ID : returned Title must be the same with input Title")
			assert.Equal(t, mockInput.Body, post.Body, "Get Post by ID : returned Body must be the same with input Body")
		})

		t.Run("TEST_UPDATE_POST_BY_ID", func(t *testing.T) {
			updateInput := models.PostInputForm{Title: "UPDATE title", Body: "This is a body of UPDATE Post"}

			updated, err := services.PostUpdate(strconv.Itoa(int(postId)), updateInput)

			assert.Equal(t, err, nil, "Update Post by ID : Must not return Error")
			assert.Equal(t, updateInput.Title, updated.Title, "Update Post by ID : updated Title must be the same with new input Title")
			assert.Equal(t, updateInput.Body, updated.Body, "Update Post by ID : updated Body must be the same with new input Body")
		})

		t.Run("TEST_DELETE_POST_BY_ID", func(t *testing.T) {
			err := services.PostDelete(strconv.Itoa(int(postId)))

			assert.Equal(t, err, nil, "Delete Post by ID : Must not return Error")
		})
	}
}
