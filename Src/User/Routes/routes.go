package user_routes

import (
	middleware "gin-template/Middleware"
	controllers "gin-template/Src/User/Controllers"
	"gin-template/Src/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	SetupAuthRoutes(r)
	User_routes := r.RouterGroup.Group("/users")
	User_routes.GET("/details", middleware.ValidateTokenByRole([]utils.Role{utils.User}), controllers.GetUserDetails)
}
