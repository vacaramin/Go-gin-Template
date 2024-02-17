package routes

import (
	middleware "ProudFlowers-Backend/Middleware"
	user_routes "ProudFlowers-Backend/Src/User/Routes"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	//Setting Up CORS Middleware
	r.Use(middleware.CorsMiddleware())

	// All Routes will be declared here
	// Call all the main routes from the source here!
	pingPongRoutes(r)
	user_routes.SetupRoutes(r)
}

// pingPongRoutes Function to set the Ping Pong GET routes on the server
func pingPongRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		c.JSON(200, gin.H{"Message": "Pong"})
	})
	r.GET("/pong", func(c *gin.Context) {
		c.Writer.WriteHeader(200)
		c.JSON(200, gin.H{"Message": "Ping"})
	})

}
