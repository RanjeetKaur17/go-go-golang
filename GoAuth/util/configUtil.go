package util

import (
	"go-go-golang/GoAuth/model"
	"os"
	"encoding/json"
)

var configuration model.Configuration

func LoadConfigurations(configPath string) {
	println("Loading Configuration")
	//open config file
	file, err := os.Open(configPath)
	if err != nil {
		println("Loading Configuration Failed" + err.Error())
	}
	decoder := json.NewDecoder(file)
	configuration = model.Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		println("Loading Configuration Failed : " + err.Error())
	} else {
		println("Loading Configuration : Success")
	}
}

func GetConfiguration() model.Configuration {
	return configuration;
}

