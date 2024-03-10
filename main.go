package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.Run(":8080")
}

// A handler method receives a pointer to the Gin Context
func getEvents(context *gin.Context) {
	// The context can be used to send a response.
	// JSON responses require a status code and an object to be converted to JSON (a map in this case).
	context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}
