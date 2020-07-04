package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	conf *Config
)

func init() {
	config = LoadConfig("config.json")
}

func main() {

	discord, err := discordgo.New(conf.DISCORD_TOKEN)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	if conf.UseSharding {
		discord.ShardID = conf.ShardId
		discord.ShardCount = conf.ShardCount
	}
	usr, err := discord.User("@me")
	if err != nil {
		fmt.Println("Error obtaining account details,", err)
		return
	}
	botId = usr.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		discord.UpdateStatus(0, conf.DefaultStatus)
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
