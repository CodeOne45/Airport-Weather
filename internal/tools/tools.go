package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"util"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// TODO : Fix the pressure function --> render real values
func PressureNumberGenerator() int {
	return rand.Intn(util.MAXIMUM_VALUE_PRESSURE-util.MINIMUM_VALUE_PRESSURE) + util.MINIMUM_VALUE_PRESSURE
}

// TODO : Fix the temp function --> render real values
func TempNumberGenerator() int {
	return rand.Intn(util.MAXIMUM_VALUE_TEMP-util.MINIMUM_VALUE_TEMP) + util.MINIMUM_VALUE_TEMP
}

// TODO : Fix the wind function --> render real values
func WindNumberGenerator() int {
	return rand.Intn(util.MAXIMUM_VALUE_WIND-util.MINIMUM_VALUE_WIND) + util.MINIMUM_VALUE_WIND
}

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