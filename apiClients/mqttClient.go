package clients

import (
	"encoding/json"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"homegear/services/device"
	"os"
	"time"
)

var Client mqtt.Client

type mqttMessageData struct {
	State bool `json:"state"`
	Id    int  `json:"id"`
}
type mqttMessage struct {
	Data    mqttMessageData `json:"data"`
	Pattern string          `json:"pattern"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

func MqttInit() {
	mqttAddress, exists := os.LookupEnv("MQTT_ADDRESS")
	if !exists {
		fmt.Println(exists)
	}

	mqttUser, exists := os.LookupEnv("MQTT_USER")
	if !exists {
		fmt.Println(exists)
	}
	mqttPassword, exists := os.LookupEnv("MQTT_PASSWORD")
	if !exists {
		fmt.Println(exists)
	}
	mqttClient, exists := os.LookupEnv("MQTT_CLIENT")
	if !exists {
		fmt.Println(exists)
	}

	opts := mqtt.NewClientOptions().AddBroker(mqttAddress).SetClientID(mqttClient)
	opts.SetProtocolVersion(4)
	opts.SetKeepAlive(4 * time.Second)
	opts.SetPingTimeout(2 * time.Second)
	username := flag.String("username", mqttUser, "A username to authenticate to the MQTT server")
	password := flag.String("password", mqttPassword, "Password to match username")
	opts.SetUsername(*username)
	opts.SetPassword(*password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	Client := mqtt.NewClient(opts)

	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := Client.Subscribe("Devices/#", 1, device.HandleReceivedMessage); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())

		os.Exit(1)
	}
	state := mqttMessage{
		Data: mqttMessageData{
			State: true,
			Id:    1,
		},
		Pattern: "Devices/Bulb/state",
	}

	payload, err := json.Marshal(state)
	if err != nil {
		fmt.Println(err)
	}
	token := Client.Publish("Devices/Bulb/state", 1, false, payload)
	token.Wait()
}
