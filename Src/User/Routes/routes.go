package user_routes

import (
	middleware "ProudFlowers-Backend/Middleware"
	controllers "ProudFlowers-Backend/Src/User/Controllers"
	"ProudFlowers-Backend/Src/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	SetupAuthRoutes(r)
	User_routes := r.RouterGroup.Group("/users")
	User_routes.GET("/details", middleware.ValidateTokenByRole([]utils.Role{utils.User}), controllers.GetUserDetails)
}
