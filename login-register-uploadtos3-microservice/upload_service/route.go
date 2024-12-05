package main

import (
	"middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/images", middlewares.Authenticate, GetUserImages)
	server.POST("/image", middlewares.Authenticate, UploadImage)
	server.GET("/s3-buckets", middlewares.Authenticate, ListS3Buckets)
}
