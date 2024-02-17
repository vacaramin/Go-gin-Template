package user_routes

import (
	controllers "ProudFlowers-Backend/Src/User/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.POST("/token-refresh", controllers.RefreshToken)

}
