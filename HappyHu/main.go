package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bwmarrin/discordgo"
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
}

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err)
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err)
	}
	BotID = user.ID

	goBot.AddHandler(messageHandler)

	//unfinished
}
