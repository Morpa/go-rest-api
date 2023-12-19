package main

import (
	"github.com/Morpa/go-rest-api/db"
	"github.com/Morpa/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
