package router

import (
	"github.com/Xuduoteng/gomall/internal/controllers"
	"github.com/Xuduoteng/gomall/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var exampleController = new(controllers.ExampleController)

func LoadExampleRoutes(r *gin.Engine) *gin.RouterGroup {

	example := r.Group("/examples")
	example.Use(middlewares.Jwt())

	{
		example.POST("/createExample", exampleController.CreateExample)
		example.GET("/getExample", exampleController.GetExample)
		example.POST("/updateExample", exampleController.UpdateExample)
		example.POST("/deleteExample", exampleController.DeleteExample)
	}
	return example
}
