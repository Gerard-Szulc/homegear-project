package usersController

import (
	"github.com/gin-gonic/gin"
	users "homegear/services/user"
)

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	users.GetUser(userID, c)
}

func GetUsers(c *gin.Context) {
	users.GetUsers(c)
}
