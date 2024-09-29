package controllers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"homegear/controllers/authenticationController"
	"homegear/controllers/devicesController"
	"homegear/controllers/measurementsController"
	"homegear/controllers/usersController"
	"homegear/docs"
	"homegear/middlewares"
)

// @BasePath /api/v1

func HandleRequests(router *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/"
	docs.SwaggerInfo.Host = "localhost:7684"

	auth := router.Group("/api/")
	{
		auth.POST("login", authenticationController.Login)
		auth.POST("register", authenticationController.Register)
	}

	authenticated := router.Group("/api/")
	authenticated.Use(middlewares.AuthenticationRequired())
	{
		authenticated.POST("login-device", authenticationController.LoginDevice)
		authenticated.POST("measurement", measurementsController.PostMeasurement)
		authenticated.POST("measurements", measurementsController.PostMeasurements)
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
	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
