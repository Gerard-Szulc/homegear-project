package authenticationController

import (
	"github.com/gin-gonic/gin"
	"homegear/services/authentication"
	"homegear/structs"
	"net/http"
)

// Login godoc
// @Summary Login user using username and password
// @Schemes
// @Description Login user using username and password
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {object} structs.LoginDto
// @Router /api/login [post]
func Login(c *gin.Context) {
	var formattedBody structs.LoginDto
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authentication.Login(formattedBody.Username, formattedBody.Password, c)
}

// LoginDevice godoc
// @Summary Login device using deviceId and token
// @Schemes
// @Description Login device using deviceId and token
// @Tags authentication
// @Body {object} structs.LoginDeviceDto
// @Accept json
// @Produce json
// @Success 200 {object} structs.LoginDeviceResponseViewModel
// @Router /api/login-device [post]
func LoginDevice(c *gin.Context) {
	var formattedBody structs.LoginDeviceDto
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authentication.LoginDevice(formattedBody.DeviceId, formattedBody.Token, c)
}

// Register godoc
// @Summary Register a new user
// @Schemes
// @Description Register a new user with username, email, and password
// @Tags authentication
// @Accept json
// @Produce json
// @Success 200 {object} structs.ResponseUserWithToken
// @Router /api/register [post]
func Register(c *gin.Context) {
	var formattedBody structs.Register
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	authentication.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password, c)
}
