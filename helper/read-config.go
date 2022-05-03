package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Username   string `json:"username"`
	LangsCount int    `json:"langs_count"`
	HideBorder bool   `json:"hide_border"`
}

func ReadConfig() Config {

	jsonFile, err := os.Open("config.json")

	if err != nil {
		log.Fatalln(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}

	var config Config

	json.Unmarshal(byteValue, &config)

	return config
}
