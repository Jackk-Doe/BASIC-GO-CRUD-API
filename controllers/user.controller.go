package controllers

import (
	"net/http"

	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserSignUp(c *gin.Context) {
	var signUpForm models.UserSignUpForm
	if bindingErr := c.BindJSON(&signUpForm); bindingErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingErr.Error()})
		return
	}

	// Check Email already existed
	_, isEmailExisted, emailCheckErr := services.UserGetViaEmail(signUpForm.Email)
	if emailCheckErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": emailCheckErr.Error()})
		return
	}
	if isEmailExisted == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This is Email is already used to register"})
		return
	}

	// Check Name already existed
	_, isNameExisted, emailCheckErr := services.UserGetViaEmail(signUpForm.Email)
	if emailCheckErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": emailCheckErr.Error()})
		return
	}
	if isNameExisted == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This is Name is already taken"})
		return
	}

	// Hash input Password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(signUpForm.Password), 10)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to hash the input Password"})
		return
	}

	createdUser, createErr := services.UserCreate(signUpForm.Email, signUpForm.Name, string(hashedPassword))
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to create User into database"})
		return
	}

	// Create JWT Token
	token, tokenErr := services.TokenCreate(createdUser.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	// Convert to UserDTO, before sending
	userDto := createdUser.ToUserDTO(token)

	// TODO : return
	c.JSON(http.StatusCreated, gin.H{"user": userDto})
}
