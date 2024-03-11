package controllers

import (
	"github.com/gin-gonic/gin"
	"homegear/controllers/authenticationController"
	"homegear/controllers/devicesController"
	"homegear/controllers/measurementsController"
	"homegear/controllers/usersController"
	"homegear/middlewares"
)

func HandleRequests(router *gin.Engine) {

	auth := router.Group("/api/")
	{
		auth.POST("login", authenticationController.Login)
		auth.POST("register", authenticationController.Register)
	}

	authenticated := router.Group("/api/")
	authenticated.Use(middlewares.AuthenticationRequired())
	{
		authenticated.POST("login-device", authenticationController.LoginDevice)
		//testing pi pico and esp
		authenticated.POST("measurements", measurementsController.PostMeasurement)
		authenticated.GET("measurements/:deviceId", measurementsController.GetMeasurements)

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
