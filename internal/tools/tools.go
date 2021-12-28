package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func ReadFile(filename string) []byte {
	jsonfile, err := os.Open("../../internal/util/" + filename + ".json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonfile.Close()
	bytevalue, err := ioutil.ReadFile(jsonfile.Name())
	if err != nil {
		fmt.Println(err)
	}
	return bytevalue
}

func CreateClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientId)
	return opts
}

func Connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := CreateClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	// Get client token
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}