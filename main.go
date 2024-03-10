package main

import (
	"github.com/amanjshah/event-booking-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")
}

// A handler function receives a pointer to the Gin Context
func getEvents(context *gin.Context) {
	models.GetAllEvents()
	// The context can be used to send a response.
	// JSON responses require a status code and an object to be converted to JSON.
	context.JSON(http.StatusOK, models.GetAllEvents())
}

func createEvent(context *gin.Context) {
	var event models.Event
	// ShouldBindJson binds the request body to a pointer input var.
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	// gin.H is an alias for map[string]any
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
