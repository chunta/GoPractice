package main

import (
	"restful/db"
	"restful/routes"

	"github.com/gin-gonic/gin"
)

var autoId = 0

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8110")
}
