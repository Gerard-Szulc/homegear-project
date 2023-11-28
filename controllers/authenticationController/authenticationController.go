package authenticationController

import (
	"github.com/gin-gonic/gin"
	"homegear/services/authentication"
	"homegear/structs"
	"net/http"
)

func Login(c *gin.Context) {
	var formattedBody structs.LoginDto
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authentication.Login(formattedBody.Username, formattedBody.Password, c)
}

func LoginDevice(c *gin.Context) {
	var formattedBody structs.LoginDeviceDto
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authentication.LoginDevice(formattedBody.DeviceId, formattedBody.Token, c)
}

func Register(c *gin.Context) {
	var formattedBody structs.Register
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authentication.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password, c)
}
