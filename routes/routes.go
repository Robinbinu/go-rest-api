package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", rootHandler)
	server.GET("/events", getEvents)
	// :id dynamic param in gin id can be anything 1/2/3/etc
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id",updateEvent)
}
