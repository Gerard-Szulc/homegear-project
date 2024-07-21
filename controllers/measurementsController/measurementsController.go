package measurementsController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homegear/db"
	measurementsService "homegear/services/measurements"
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

func PostMeasurements(c *gin.Context) {
	var measurementData structs.MeasurementsDto
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

	measurements := structs.Measurements{}

	for _, data := range measurementData {
		singleMeasurement := structs.Measurement{
			Type:   data.Type,
			Value:  data.Value,
			Device: *device,
		}
		measurements = append(measurements, singleMeasurement)
	}
	results := db.DB.Create(&measurements)
	fmt.Println(results.Error)

	c.JSON(http.StatusOK, measurements)
}

func GetMeasurements(c *gin.Context) {
	deviceId := c.Param("deviceId")
	if deviceId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "deviceId not found"})
		return
	}

	userId, _ := c.Get("userId")

	if measurementsService.UserOwnsDevice(userId.(float64), deviceId) == false {
		c.JSON(http.StatusForbidden, gin.H{"error": "device not found"})
		return
	}
	fmt.Println("after")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	measurements, _ := measurementsService.GetMeasurements(deviceId, startDate, endDate)
	c.JSON(http.StatusOK, measurements)
}
