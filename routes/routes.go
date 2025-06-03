package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", rootHandler)
	server.GET("/events", getEvents)
	// :id dynamic param in gin id can be anything 1/2/3/etc
	server.GET("/events/:id", getEvent)
	server.POST("/signup",signup)
	server.POST("/login",login)

	//a group for which a middleware can be applied as a whole
	//route protection is enabled for this group of routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)
}
