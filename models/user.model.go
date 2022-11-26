package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// For interaction with database only
type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Name     string `gorm:"unique"`
	Password string
}

// For User Sign-Up form receiving
type UserSignUpForm struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,max=30"`
	Password string `json:"password" binding:"required,min=8"`
}

// For User Log-In form receiving
type UserLogInForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// For sending to backend caller
type UserDTO struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Token string `json:"jwt" binding:"required"`
}

// This function is run before [Post] is created into database
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	idString := uuid.New().String()
	u.ID = idString
	return
}

// To convert from [User] to [UserDTO], with input token string
func (u *User) ToUserDTO(token string) UserDTO {
	return UserDTO{Email: u.Email, Name: u.Name, Token: token}
}
