package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

func OpenConfig() {
	fmt.Println("Reading config file...")
	configFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Config opened successfully")
	defer configFile.Close()
	byteValue, _ := io.ReadAll(configFile)

	json.Unmarshal(byteValue, &config)

	Token = config.Token
	BotPrefix = config.BotPrefix
	fmt.Println("Bot Token" + config.Token)
	fmt.Println("Bot Prefix: " + config.BotPrefix)

	return
}
