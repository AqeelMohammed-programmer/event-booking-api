package main

import (
	"net/http"

	"github.com/AqeelMohammed-programmer/event-booking-api/db"
	"github.com/AqeelMohammed-programmer/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot fetch the events, try again later!",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse requested date",
		})
	}
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create an event, try again leter!",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created!", "event": event,
	})
}
