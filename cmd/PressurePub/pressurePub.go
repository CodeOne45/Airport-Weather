package main

import (
	"encoding/json"
	"tools"
	"util"
)

func main(){
	// Use pressure configuration
	values := tools.ReadFile("pressure-config")
	var config util.Config
	json.Unmarshal(values, &config)

	// Creation/Connection of a client
	client := tools.Connect(config.Host, config.ClientID)

	for{
		//client.Publish
	}
}