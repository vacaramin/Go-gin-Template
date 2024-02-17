package main

import (
	initializers "gin-template/Initializers"
	middleware "gin-template/Middleware"
	routes "gin-template/Routes"
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
