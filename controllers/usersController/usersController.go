package usersController

import (
	users "dustData/services/user"
	"dustData/structs"
	"dustData/utils"
	"github.com/gin-gonic/gin"
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

	//auth := r.Header.Get("Authorization")
	if !utils.ValidateRequestToken(c.Request) {
		c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{
			"message": "error:token_not_valid",
		}})
	}
	users.GetUser(userID, c)
}

func GetUsers(c *gin.Context) {
	if !utils.ValidateRequestToken(c.Request) {
		c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error:token_not_valid"}})
		return
	}
	users.GetUsers(c)
}
