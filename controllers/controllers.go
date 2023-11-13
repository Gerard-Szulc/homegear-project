package controllers

import (
	"dustData/controllers/devicesController"
	"dustData/controllers/measurementsController"
	"dustData/controllers/usersController"
	"dustData/middlewares"
	"github.com/gin-gonic/gin"
)

func HandleRequests(router *gin.Engine) {

	auth := router.Group("/api/")
	{

		auth.POST("login", usersController.Login)
		auth.POST("register", usersController.Register)

		//testing pi pico and esp without provisioning
		auth.POST("measurements", measurementsController.PostMeasurement)
	}

	authenticated := router.Group("/api/")
	authenticated.Use(middlewares.AuthenticationRequired())
	{
		users := authenticated.Group("/users")
		{
			users.GET("/:id", usersController.GetUser)
		}

		devices := authenticated.Group("/devices")
		{
			devices.POST("/", devicesController.PostDevice)
			devices.GET("/", devicesController.GetDevices)
			devices.GET("/:id", devicesController.GetDevice)
		}
	}
}
