package services

import (
	"time"

	"github.com/Jackk-Doe/basic-go-crud-api/database"
	"github.com/Jackk-Doe/basic-go-crud-api/models"
	"github.com/Jackk-Doe/basic-go-crud-api/shared"
	"github.com/golang-jwt/jwt/v4"
)

// Find User via input email; if found return : User, bool: True, err: nil
func UserGetViaEmail(email string) (models.User, bool, error) {
	var user models.User
	dbIns := database.GetDB()
	result := dbIns.Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return models.User{}, false, result.Error //! DB error
	}
	if result.RowsAffected < 1 {
		return models.User{}, false, nil //NOT FOUND
	}
	return user, true, nil //FOUND
}

// Find User via input name; if found return : User, bool: True, err: nil
func UserGetViaName(name string) (models.User, bool, error) {
	var user models.User
	dbIns := database.GetDB()
	result := dbIns.Where("name = ?", name).Find(&user)
	if result.Error != nil {
		return models.User{}, false, result.Error //! DB error
	}
	if result.RowsAffected < 1 {
		return models.User{}, false, nil //NOT FOUND
	}
	return user, true, nil //FOUND
}

func UserCreate(email, name, password string) (models.User, error) {
	user := models.User{Email: email, Name: name, Password: password}
	dbIns := database.GetDB()
	err := dbIns.Create(&user).Error
	return user, err
}

// Create token from User.ID and signing with Token Secret Key
func TokenCreate(userId string) (string, error) {
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(), //Expiration : 24hour * 30days
	})

	tokenSecretKey := shared.GetTokenSecretKey()
	convertedSecretKey := []byte(tokenSecretKey)

	// Sign off to get the complete encoded token, encoded with the Token Secret Key
	signedToken, tokenSigningErr := unsignedToken.SignedString(convertedSecretKey)
	return signedToken, tokenSigningErr
}
