package main

import (
	"./old/commands"
	"./old/domains"
	"errors"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	sessions   *domains.SessionManager
	cmdHandler *domains.CommandHandler
	config     *domains.Config
	//botID      string
)

func init() {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		panic(errors.New("CONFIG_FILE is not defined"))
	}
	config = domains.LoadConfig(configFile)
}

func main() {
	cmdHandler = domains.NewCommandHandler()
	sessions = domains.NewSessionManager()
	// Load all commands into the bot
	registerAllCommands()

	// Create a discord session
	log.Info("Starting discord session...")
	discord, err := discordgo.New(config.BotToken)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}

	if config.UseSharding {
		discord.ShardID = config.ShardId
		discord.ShardCount = config.ShardCount
	}

	discord.AddHandler(commandHandler)

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
	defer discord.Close()

	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Info("Closing sessions.")
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	content := message.Content
	if strings.HasPrefix(content, config.Prefix) {
		user := message.Author
		if user.Bot {
			return
		}
		// remove the Prefix from content
		content = content[len(config.Prefix):]
		if len(content) < 1 {
			return
		}
		args := strings.Fields(content)
		name := strings.ToLower(args[0])

		command, found := cmdHandler.Get(name)
		if !found {
			return
		}
		channel, err := discord.State.Channel(message.ChannelID)
		if err != nil {
			fmt.Println("Error getting channel,", err)
			return
		}
		guild, err := discord.State.Guild(channel.GuildID)
		if err != nil {
			fmt.Println("Error getting guild,", err)
			return
		}
		ctx := domains.NewContext(discord, guild, channel, user, message, config, cmdHandler, sessions)
		ctx.Args = args[1:]
		c := *command
		c(*ctx)
	}

}

func registerAllCommands() {
	log.Info("Registering all commands")
	cmdHandler.Register("meme", commands.MemeCommand, "LA FETE")

	log.Debug("Loading Memes")
	memes := domains.LoadMemes(os.Getenv("MEME_FILE"))
	log.WithFields(log.Fields{
		"memes": memes,
	}).Debug("Memes loaded")
	for index, meme := range *memes {
		cmdHandler.RegisterMemeCmd(meme.Command, commands.MemeCommand, meme.Help, meme.YoutubeURL, meme.FileName)
		log.WithFields(log.Fields{
			"index": index,
			"memes": meme.Command,
		}).Debug("Registering meme command")
	}
}
