package main

import (
	"example.com/event-booking-api/db"
	"example.com/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDatabase()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run("127.0.0.1:8080")
}
