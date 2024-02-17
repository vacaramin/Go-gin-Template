package controllers

import (
	initializers "ProudFlowers-Backend/Initializers"
	model "ProudFlowers-Backend/Src/User/Model"
	"ProudFlowers-Backend/Src/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// Get the email/pass off req Body
	var body struct {
		Email    string
		Password string
		Username string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password.",
		})
		return
	}

	// Create the user
	user := model.User{Email: body.Email, Password: string(hash), Username: body.Username}

	if initializers.DB.First(&user, "email = ?", body.Email).RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user.",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User created successfully.",
		"user":    user,
	})
}

func Login(c *gin.Context) {
	// Get email & pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up for requested user
	var user model.User

	initializers.DB.First(&user, "Email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in password with saved users password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()
	rt, err := refreshToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create Refreshtoken",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access-token":  tokenString,
		"refresh-token": rt,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	// user.(models.User).Email    -->   to access specific data

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func Logout(c *gin.Context) {

	user, err := utils.ReturnUserFromToken(c)
	log.Println(user)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Logout Successful",
		})
	}
}
func RefreshToken(c *gin.Context) {
	// Define a struct to hold the request body
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}
	// Bind the JSON request body to the defined struct
	err := c.BindJSON(&body)
	if err != nil {
		log.Println("JSON Binding failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process request"})
		return
	}

	// Parse the refresh token from the request body
	token, err := jwt.Parse(body.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Provide the secret key used for signing the token
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
		return
	}
	// Check if token parsing was successful and the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the subject of the token matches the expected value
		userID, ok := claims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token format"})
			return
		}

		// Generate a new access token with user-specific information
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": userID,
			"exp": time.Now().Add(time.Hour * 2).Unix(), // You can adjust the expiration time as needed
		})
		// Sign and get the complete encoded token as a string using the secret
		newAccessToken, err := newToken.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
			return
		}
		// Return the new access token as JSON response
		c.JSON(http.StatusOK, gin.H{"access_token": newAccessToken})
		return
	}
	// If token parsing or validation fails, return error response
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
