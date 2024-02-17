package utils

import (
	"errors"
	initializers "gin-template/Initializers"
	model "gin-template/Src/User/Model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// ReturnUserFromToken This function is utility function that returns the details of the user based on the Token
func ReturnUserFromToken(c *gin.Context) (*model.User, error) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, errors.New("invalid token")
	}
	// Check if the token is expired
	expTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expTime) {
		return nil, errors.New("token has expired")
	}
	userSub := int(claims["sub"].(float64))

	// Fetch the user from the database using user email and attach it to the context
	var user model.User
	initializers.DB.First(&user, userSub)
	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return nil, err
	}
	return &user, nil

}
