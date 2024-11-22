package router

import (
	"github.com/Xuduoteng/gomall/internal/controllers"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		user.POST("/register", userController.CreateUser)
		user.POST("/login", userController.LoginByUsernamePassword)
	}
	return user
}
