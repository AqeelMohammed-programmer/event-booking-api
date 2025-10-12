package routes

import (
	"github.com/AqeelMohammed-programmer/event-booking-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
	server.DELETE("/events/:id", controllers.DeleteEvent)
	server.POST("/signup", controllers.SignUp)
}
