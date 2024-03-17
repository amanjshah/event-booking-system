package routes

import (
	"github.com/amanjshah/event-booking-system/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
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

	err = event.Register(context.GetInt64("userId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user for event. Please retry."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered! "})
}

func cancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("eventId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(context.GetInt64("userId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel registration of user for event. Please retry."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled! "})
}
