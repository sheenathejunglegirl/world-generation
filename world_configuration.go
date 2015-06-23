package main

import (
	"encoding/json"
	"os"
)

type WorldConfiguration struct {
	Map MapConfiguration
}

type MapConfiguration struct {
	Min int
	Max int
}

func getWorldConfiguration(fileName string) (WorldConfiguration, error) {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := WorldConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}
