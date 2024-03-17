package routes

import (
	"github.com/amanjshah/event-booking-system/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventId", updateEvent)
	authenticated.DELETE("/events/:eventId", deleteEvent)
}
