package measurementsController

import (
	"dustData/db"
	"dustData/structs"
	"fmt"
	"github.com/gin-gonic/gin"
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

	// testing phase
	db.DB.Where("id = ? ", 3).First(&device)

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
