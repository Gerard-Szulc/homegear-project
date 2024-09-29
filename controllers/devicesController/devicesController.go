package devicesController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homegear/db"
	"homegear/structs"
	"net/http"
)

// PostDevice godoc
// @Summary Add a new device
// @Schemes
// @Description Add a new device with name, label, type, and IP
// @Tags devices
// @Body {object} structs.DeviceDto
// @Accept json
// @Produce json
// @Success 200 {object} structs.DeviceResponse
// @Router /api/devices [post]
func PostDevice(c *gin.Context) {
	var deviceData structs.DeviceDto
	if err := c.BindJSON(&deviceData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user structs.User
	userId, _ := c.Get("userId")
	db.DB.Where("id = ? ", userId).First(&user)
	users := []*structs.User{&user}

	device := structs.Device{
		Name:  deviceData.Name,
		Label: deviceData.Label,
		Type:  deviceData.Type,
		IP:    deviceData.IP,
		Users: users,
	}

	results := db.DB.Create(&device)
	if results.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error})
		return
	}

	fmt.Println(results.Error)
	deviceResponse := structs.DeviceResponse{
		Id:    int(device.ID),
		Name:  device.Name,
		Label: device.Label,
		Type:  device.Type,
		IP:    device.IP,
	}

	c.JSON(http.StatusOK, deviceResponse)
}

// GetDevices godoc
// @Summary Get all devices
// @Schemes
// @Description Retrieve a list of all devices
// @Tags devices
// @Produce json
// @Success 200 {array} structs.DeviceDto
// @Router /api/devices [get]
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
