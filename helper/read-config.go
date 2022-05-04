package helper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Wakatime Wakatime `json:"wakatime"`
	Github   Github   `json:"github"`
}

type Wakatime struct {
	Username   string `json:"username"`
	LangsCount int    `json:"langs_count"`
	HideBorder bool   `json:"hide_border"`
}

type Github struct {
	Username     string `json:"username"`
	ShowIcons    bool   `json:"show_icons"`
	HideBorder   bool   `json:"hide_border"`
	CountPrivate bool   `json:"count_private"`
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
