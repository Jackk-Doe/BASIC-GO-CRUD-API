package services_test

import (
	"log"
	"testing"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/services"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_User_Services_Functions(t *testing.T) {
	log.Println()
	log.Println("--> START : User Services unit tests...")
	log.Println()

	/// Use this hashed password for Mock User's password
	nonHashedPassword := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(nonHashedPassword), 10)
	mockPassword := string(hashedPassword)

	/// Mock User data Sign-Up Form
	mockUser := models.UserSignUpForm{Email: "mock-user@email.com", Name: "MockUser", Password: mockPassword}

	// Use this to store created User ID
	userId := ""

	t.Run("TEST_CREATE_USER", func(t *testing.T) {
		createdUser, err := services.UserCreate(mockUser.Email, mockUser.Name, mockUser.Password)

		// Assign created User ID
		userId = createdUser.ID

		assert.Nil(t, err, "Create User function : Must not return Error")
		assert.Equal(t, mockUser.Email, createdUser.Email, "Create User function : created User Email must be same with input Email")
		assert.Equal(t, mockUser.Name, createdUser.Name, "Create User function : created User Name must be same with input Name")
		assert.NotNil(t, createdUser.ID, "Create User function : created User ID must be not be Nil")
		assert.Nil(t, createdUser.PasswordMatchCheck(nonHashedPassword), "Create User function : created User password-check-validation must ok, with given input Password")
	})

	t.Run("TEST_GET_USER_VIA_EMAIL", func(t *testing.T) {
		user, isFound, err := services.UserGetViaEmail(mockUser.Email)
		assert.Nil(t, err, "Get User via Email function : Must not return Error")
		assert.True(t, isFound, "Get User via Email function : User must be found")
		assert.Equal(t, mockUser.Email, user.Email, "Get User via Email function : User Email must match with input Email")
		assert.Equal(t, mockUser.Name, user.Name, "Get User via Email function : User Name must match with input Name")
		assert.Nil(t, user.PasswordMatchCheck(nonHashedPassword), "Get User via Email function : User password-check-validation must ok, with given input Password")
	})

	t.Run("TEST_GET_USER_VIA_NAME", func(t *testing.T) {
		user, isFound, err := services.UserGetViaName(mockUser.Name)
		assert.Nil(t, err, "Get User via Name function : Must not return Error")
		assert.True(t, isFound, "Get User via Name function : User must be found")
		assert.Equal(t, mockUser.Email, user.Email, "Get User via Name function : User Email must match with input Email")
		assert.Equal(t, mockUser.Name, user.Name, "Get User via Name function : User Name must match with input Name")
		assert.Nil(t, user.PasswordMatchCheck(nonHashedPassword), "Get User via Name function : User password-check-validation must ok, with given input Password")
	})

	t.Run("TEST_GET_USER_VIA_ID", func(t *testing.T) {
		user, isFound, err := services.UserGetViaID(userId)
		assert.Nil(t, err, "Get User via ID function : Must not return Error")
		assert.True(t, isFound, "Get User via ID function : User must be found")
		assert.Equal(t, mockUser.Email, user.Email, "Get User via ID function : User Email must match with input Email")
		assert.Equal(t, mockUser.Name, user.Name, "Get User via ID function : User Name must match with input Name")
		assert.Nil(t, user.PasswordMatchCheck(nonHashedPassword), "Get User via ID function : User password-check-validation must ok, with given input Password")
	})

	t.Run("TEST_CREATE_JWT_TOKEN_FROM_USER_ID", func(t *testing.T) {
		token, err := services.TokenCreate(userId)
		assert.Nil(t, err, "Create Token function : Must not return Error")
		assert.NotEmpty(t, token, "Create Token function : created Token must not be empty")
	})

	log.Println()
	log.Println("--> END : User Services unit tests...")
	log.Println()
}
