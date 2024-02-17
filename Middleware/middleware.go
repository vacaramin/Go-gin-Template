package middleware

import (
	initializers "ProudFlowers-Backend/Initializers"
	model "ProudFlowers-Backend/Src/User/Model"
	"ProudFlowers-Backend/Src/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ValidateTokenByRole(roles []utils.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Message": "Please Provide Auth Token with Request",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Invalid Token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "You are Unauthorized",
			})
			c.Abort()
			return
		}

		userSub := int(claims["sub"].(float64))

		// Fetch the user from the database using user ID
		var user model.User

		initializers.DB.First(&user, userSub)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "User Not found",
			})
			c.Abort()
			return
		}

		// Check if the user's role matches any of the specified roles
		roleMatch := false
		for _, role := range roles {
			if user.Role == string(role) {
				roleMatch = true
				break
			}
		}

		if !roleMatch {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "You are Not Authorized to make this request",
			})
			c.Abort()
			return
		}

		// User is authorized, proceed with the next middleware or handler
		c.Next()
	}
}
