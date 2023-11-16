package usersController

import (
	"github.com/gin-gonic/gin"
	users "homegear/services/user"
	"homegear/structs"
	"homegear/utils"
	"io"
	"net/http"
)

type LoginDto struct {
	Username string
	Password string
}

func readBody(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	utils.HandleErr(err)

	return body
}

func Login(c *gin.Context) {

	var formattedBody LoginDto
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users.Login(formattedBody.Username, formattedBody.Password, c)
}

func Register(c *gin.Context) {
	var formattedBody structs.Register
	if err := c.BindJSON(&formattedBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//utils.HandleErr(err)
	users.Register(formattedBody.Username, formattedBody.Email, formattedBody.Password, c)
}

func GetUser(c *gin.Context) {
	userID := c.Param("id")
	users.GetUser(userID, c)
}

func GetUsers(c *gin.Context) {
	users.GetUsers(c)
}
