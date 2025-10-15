package routes

import (
	"github.com/AqeelMohammed-programmer/event-booking-api/controllers"
	"github.com/AqeelMohammed-programmer/event-booking-api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	authenticated.POST("/events/:id/register", controllers.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", controllers.RegisterForEvent)

	server.POST("/signup", controllers.SignUp)
	server.POST("/login", controllers.Login)
}
