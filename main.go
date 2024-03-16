package main

import (
	"github.com/amanjshah/event-booking-system/db"
	"github.com/amanjshah/event-booking-system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
