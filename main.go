package main

import (
	"dustData/controllers"
	"dustData/db"
	"dustData/structs"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitiateDB()
	db.DB.AutoMigrate(&structs.Device{})
	db.DB.AutoMigrate(&structs.User{})
	db.DB.AutoMigrate(&structs.Measurement{})

	//db.InitiateMongo()

	// mqtt init
	//clients.MqttInit()

	r := gin.Default()

	// prepare frontend later
	//r.Static("/assets", "dist/assets")
	//r.LoadHTMLGlob("dist/*.html")
	//r.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", nil)
	//})

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	controllers.HandleRequests(r)
	r.Run()
}
