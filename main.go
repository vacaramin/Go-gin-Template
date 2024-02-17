package main

import (
	initializers "ProudFlowers-Backend/Initializers"
	middleware "ProudFlowers-Backend/Middleware"
	routes "ProudFlowers-Backend/Routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.Init()
	router := gin.Default()
	//Setting Up CORS Middleware
	router.Use(middleware.CorsMiddleware())

	routes.SetupRoutes(router)
	router.Run(":" + os.Getenv("PORT"))
}
