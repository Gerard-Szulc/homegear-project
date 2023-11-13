package device

import (
	"dustData/structs"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
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
