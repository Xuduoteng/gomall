package router

import (
	"github.com/Xuduoteng/gomall/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Init() {
	Router = gin.Default()

	// Global middlewares
	Router.Use(middlewares.ErrorHandle())
	Router.Use(middlewares.Cors())

	// public routes, no auth required
	LoadPublicRoutes(Router)

	// user routes
	LoadUserRoutes(Router)

	LoadProductRoutes(Router)

	// example routes
	LoadExampleRoutes(Router)

	// init swagger
	// Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
