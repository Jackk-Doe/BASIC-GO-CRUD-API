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

	// Check if input Email already existed
	_, isEmailExisted, emailCheckErr := services.UserGetViaEmail(signUpForm.Email)
	if emailCheckErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": emailCheckErr.Error()})
		return
	}
	if isEmailExisted == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This Input Email is already used to register"})
		return
	}

	// Check if input Name already existed
	_, isNameExisted, nameCheckErr := services.UserGetViaName(signUpForm.Name)
	if nameCheckErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": nameCheckErr.Error()})
		return
	}
	if isNameExisted == true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This Input Name is already taken"})
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

	c.JSON(http.StatusCreated, gin.H{"user": userDto})
}

func UserLogIn(c *gin.Context) {
	var logInForm models.UserLogInForm
	if bindingErr := c.BindJSON(&logInForm); bindingErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": bindingErr.Error()})
		return
	}

	// Check User with input Email existed
	user, isEmailExisted, emailCheckErr := services.UserGetViaEmail(logInForm.Email)
	if emailCheckErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": emailCheckErr.Error()})
		return
	}
	if isEmailExisted == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User of this Email is not yet registered"})
		return
	}

	// Check Password matching
	if passwordMatchErr := user.PasswordMatchCheck(logInForm.Password); passwordMatchErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password is not matched"})
		return
	}

	// Create JWT Token
	token, tokenErr := services.TokenCreate(user.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	// Convert to UserDTO, before sending
	userDto := user.ToUserDTO(token)

	c.JSON(http.StatusOK, gin.H{"user": userDto})
}

func GetMe(c *gin.Context) {
	userId, isFoundID := c.Get("user_id")
	if isFoundID == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Field [user_id] can not be found inside the request",
		})
		return
	}

	// NOTE : Convert form any -> string
	convertedUserId := userId.(string)

	// Get User via ID, also check User existing
	user, isFoundUser, findErr := services.UserGetViaID(convertedUserId)
	if findErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": findErr.Error(),
		})
		return
	}
	if isFoundUser == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User with the given ID is not found",
		})
		return
	}

	// Convert to UserDTONoToken, before sending
	userDto := user.ToUserDTONoToken()

	c.JSON(http.StatusOK, gin.H{"user": userDto})
}
