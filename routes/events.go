package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func rootHandler(context *gin.Context) {
	//to return a json when request is made
	context.JSON(http.StatusOK, gin.H{
		"message":  "root",
		"location": "/",
	})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	//set userId of event from the verified token data
	eUserID, ok := context.Get("userID")
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "userid not found",
		})
	}
	event.UserID = eUserID.(int64)
	err := context.ShouldBindJSON(&event)
	log.Default().Println(event, " ", context.GetInt64("userId"))
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data."})
		return
	}
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"Event":   event,
	})
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"event": event})

}

func updateEvent(context *gin.Context) {
	//get id from url
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	//check if the event exists
	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	userID := context.GetInt64("userID")
	eventUserID := event.UserID
	if eventUserID != userID {
		context.JSON(http.StatusBadRequest, gin.H{"message": "userid mismatch"})
		return
	}

	//get the updated event from request
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	updateEvent.UserID = eventID

	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "invalid event",
				"error": err})
		return
	}

	//update event in database
	updateEvent.ID = eventID
	err = updateEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "event cannot be updated",
				"error": err})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "event updated successfully",
			"event":   updateEvent})

}

func deleteEvent(context *gin.Context) {
	//get id from url
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	//check if the event exists
	event, err := models.GetEventById(eventID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	userID := context.GetInt64("userID")
	eventUserID := event.UserID
	if eventUserID != userID {
		context.JSON(http.StatusBadRequest, gin.H{"message": "userid mismatch unauthorised"})
		return
	}

	//delete event
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error while deleting event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})

}


