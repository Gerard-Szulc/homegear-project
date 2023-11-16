package devicesController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homegear/db"
	"homegear/structs"
	"net/http"
)

func PostDevice(c *gin.Context) {
	var deviceData structs.DeviceDto
	if err := c.BindJSON(&deviceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	device := structs.Device{
		Name:  deviceData.Name,
		Label: deviceData.Label,
		Type:  deviceData.Type,
		IP:    deviceData.IP,
	}

	results := db.DB.Create(&device)
	fmt.Println(results.Error)

	c.JSON(http.StatusOK, device)
}

func GetDevices(c *gin.Context) {
	var devicesData structs.DevicesDto

	db.DB.Raw("select * from devices").Scan(&devicesData)

	c.JSON(http.StatusOK, devicesData)
}

func GetDevice(c *gin.Context) {
	deviceId := c.Param("id")
	var device structs.DevicesDto

	db.DB.Where("id = ? ", deviceId).First(&device)
	c.JSON(http.StatusOK, device)
}
