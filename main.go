package main

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

var (
	config *Config
	botId  string
)

func init() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		panic(errors.New("CONFIG_FILE is not defined"))
	}
	config = LoadConfig(configFile)
}

func main() {

	discord, err := discordgo.New(config.BotToken)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	if config.UseSharding {
		discord.ShardID = config.ShardId
		discord.ShardCount = config.ShardCount
	}
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateStatus(0, config.DefaultStatus)
		guilds := discord.State.Guilds
		fmt.Println("Ready with", len(guilds), "guilds.")
	})
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}
	fmt.Println("Started")
	<-make(chan struct{})
}
