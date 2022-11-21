package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func OpenConfig() error {
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

	return nil
}

var BotID string
var goBot *discordgo.Session

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}
	if m.Content == BotPrefix+"ping" {
		message := "pong"
		_, _ = s.ChannelMessageSend(m.ChannelID, message)

	}
}

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
	fmt.Println("Bot is now running")
	// works until this part

	err = goBot.Open()

	if err != nil {
		fmt.Println(err)
	}
	goBot.Identify.Intents = discordgo.IntentsGuildMessages
}

func main() {
	err := OpenConfig()

	if err != nil {
		fmt.Println(err)
	}

	Start()

	stchan := make(chan os.Signal, 1)
	signal.Notify(stchan, syscall.SIGTERM, os.Interrupt, syscall.SIGSEGV)
end:
	for {
		select {
		case <-stchan:
			break end
		default:
		}
		time.Sleep(time.Second)
	}
}
