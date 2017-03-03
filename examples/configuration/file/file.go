package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// main is the main entry point for your app
func main() {

	type Config struct {
		Port int
	}
	config := Config{Port: 80}

	// Load the config file at this path
	configFilePath := "./config.json"
	file, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("error loading:%s %s", configFilePath, err)
	}

	// Read the data into the config struct
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("error reading:%s %s", configFilePath, err)
	}

	// Print the contents
	log.Printf("Config read:%v", config)
}
