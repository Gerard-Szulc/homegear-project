package device

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang-jwt/jwt"
	"homegear/structs"
	"homegear/utils"
	"os"
	"strings"
	"time"
)

type mqttBulbMessageData struct {
	State bool `json:"state"`
	Id    int  `json:"id"`
}
type mqttBulbMessage struct {
	Data    mqttBulbMessageData `json:"data"`
	Pattern string              `json:"pattern"`
}

func HandleBulbMessage(client mqtt.Client, message mqtt.Message) {

	var data mqttBulbMessage
	err := json.Unmarshal(message.Payload(), &data)
	if err != nil {
		fmt.Printf("Error unmarshalling message: %s\n", err)
		return
	}
	fmt.Printf("Received message on topic: %s\nId: %d\nMessage: %t\n", message.Topic(), data.Data.Id, data.Data.State)
}
func HandleReceivedMessage(client mqtt.Client, message mqtt.Message) { // Inclusion of error return
	if client == nil || message == nil {
		fmt.Printf("client or message is nil")
	}

	if strings.Contains(message.Topic(), string(structs.BulbTopic)) {
		HandleBulbMessage(client, message) // Make sure handleBulbMessage returns an error if it fails
	}
}

func prepareDeviceToken(device *structs.Device) string {
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		fmt.Println(exists)
	}
	tokenContent := jwt.MapClaims{
		"device_id": device.ID,
		"expiry":    time.Now().Add(time.Hour * 24 * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(jwtKey))
	utils.HandleErr(err)
	return token
}
