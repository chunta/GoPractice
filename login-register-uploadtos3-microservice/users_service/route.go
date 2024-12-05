package main

import (
	"middlewares"

	"github.com/gin-gonic/gin"
)

// Register the routes
func Register(server *gin.Engine) {
	server.POST("/register", RegisterUser)
	server.POST("/login", LoginUser)
	server.POST("/password", middlewares.Authenticate, ChangePassword)
	server.GET("/user/:id", middlewares.Authenticate, GetUserById)
}
