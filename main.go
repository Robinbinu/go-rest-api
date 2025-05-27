package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()
	server.GET("/",rootHandler)
	server.GET("/events",eventHandler)
	server.Run()
}

func rootHandler(context *gin.Context){
	context.JSON(http.StatusOK,gin.H{
		"message":"root",
		"location":"/",
	})
}

func eventHandler(context *gin.Context){
	context.JSON(http.StatusAccepted,"No Events Found")
}