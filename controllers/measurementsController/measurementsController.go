package measurementsController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homegear/db"
	"homegear/structs"
	"net/http"
)

func PostMeasurement(c *gin.Context) {
	var measurementData structs.MeasurementDto
	if err := c.BindJSON(&measurementData); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := &structs.Device{}
	deviceId, exists := c.Get("deviceId")
	if !exists {
		fmt.Println(exists)
		c.JSON(http.StatusBadRequest, gin.H{"error": "deviceId not found"})
		return
	}

	db.DB.Where("id = ? ", deviceId).First(&device)

	measurement := structs.Measurement{
		Type:   measurementData.Type,
		Value:  measurementData.Value,
		Device: *device,
	}

	results := db.DB.Create(&measurement)
	fmt.Println(results.Error)

	fmt.Println(measurementData.Type)
	fmt.Println(measurementData.Value)

	c.JSON(http.StatusOK, measurement)
}

func GetMeasurements(c *gin.Context) {
	deviceId, exists := c.Get("deviceId")
	if !exists {
		fmt.Println(exists)
		c.JSON(http.StatusBadRequest, gin.H{"error": "deviceId not found"})
		return
	}
	userId, _ := c.Get("userId")
	device := &structs.Device{}
	db.DB.Where("id = ? AND ? IN users ", deviceId, userId).First(&device)
	fmt.Println(device.ID)

	if device.ID == 0 {
		c.JSON(http.StatusForbidden, gin.H{"error": "device not found"})
		return
	}
	measurements := []structs.Measurement{}
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if startDate != "" && endDate != "" {
		db.DB.Where("device_id = ? AND created_at BETWEEN ? AND ?", deviceId, startDate, endDate).Find(&measurements)
	} else {
		db.DB.Where("device_id = ?", deviceId).Find(&measurements)
	}
	c.JSON(http.StatusOK, measurements)
}
