package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Jackk-Doe/basic-go-crud-api/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthViaJWT(c *gin.Context) {
	bearer_n_token := c.Request.Header.Get("Authorization")

	// The token is in splited[1]
	splited := strings.Split(bearer_n_token, " ")
	if len(splited) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token needs to be provided in the request header"})
		c.Abort()
		return
	}

	token := splited[1]

	// Decode [token] with Token Secret Key
	decodedToken, decodedErr := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := shared.GetTokenSecretKey()
		return []byte(secretKey), nil
	})

	if decodedErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Problem with Token Decoding, possibly wrong format"})
		c.Abort()
		return
	}

	// Check if the [decodedToken] is valid & has User.ID in it
	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {

		// Check Token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token is already expired"})
			c.Abort()
			return
		}

		// Try to retrieve [user_id] from the decoded token
		userId, isFoundUserId := claims["user_id"]
		if isFoundUserId == false {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to retrieve User's ID from the decoded Token"})
			c.Abort()
			return
		}

		/// Received [user_id] from claims SUCCESS,
		/// attach the [user_id] into the Request
		c.Set("user_id", userId)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Decoded token is in-valid to retrieve any datas from it"})
		c.Abort()
		return
	}

	/// All validation cleared, proceed to next process
	c.Next()
}
