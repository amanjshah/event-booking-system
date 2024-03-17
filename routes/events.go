package routes

import (
	"github.com/amanjshah/event-booking-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// A handler function receives a pointer to the Gin Context
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	// The context can be used to send a response.
	// JSON responses require a status code and an object to be converted to JSON.
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events. Please retry."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Please retry."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	// ShouldBindJson binds the request body to a pointer input var.
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.UserID = context.GetInt64("userId")

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event. Please retry."})
		return
	}
	// gin.H is an alias for map[string]any
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(context *gin.Context) {
	// Get event ID
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// Validate that an event with the specified ID currently exists
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Please retry."})
		return
	}
	if event.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, "Not authorized to update event created by different user.")
		return
	}

	// Bind request body to a new Event instance
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	// Set the ID of the new instance to the specified event ID, then update
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event. Please retry."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully! "})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event. Please retry."})
		return
	}
	if event.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, "Not authorized to delete event created by different user.")
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event. Please retry."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully! "})
}
