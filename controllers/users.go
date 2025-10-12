package controllers

import (
	"net/http"

	"github.com/AqeelMohammed-programmer/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse requested date",
		})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "cannot create a user, try again leter!",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "user created!",
	})
}
