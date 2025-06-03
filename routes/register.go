package routes

import (
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context){
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not find event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event registration unsuccessful"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event registration successful"})
}

func cancelRegistration(context *gin.Context){
	userId := context.GetInt64("userId")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	var event models.Event
	event.ID = eventID

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event registration cancelled successfully"})
}