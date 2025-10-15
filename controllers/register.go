package controllers

import (
	"net/http"
	"strconv"

	"github.com/AqeelMohammed-programmer/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the event id!",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event!",
		})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "event not found.",
		})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully!",
	})

}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse the event id!",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the event!",
		})
		return
	}

	if event == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "event not found.",
		})
		return
	}

	if !event.IsRegistered(userId) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "user not registered to this event!",
		})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "user canceled registration successfully!",
	})
}
