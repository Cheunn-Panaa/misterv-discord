package domains

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Context struct
type Context struct {
	Discord      *discordgo.Session
	Guild        *discordgo.Guild
	VoiceChannel *discordgo.Channel
	TextChannel  *discordgo.Channel
	User         *discordgo.User
	Message      *discordgo.MessageCreate
	Args         []string

	// dependency injection?
	Conf       *Config
	CmdHandler *CommandHandler
	Sessions   *SessionManager
}

// NewContext constructor
func NewContext(discord *discordgo.Session, guild *discordgo.Guild, textChannel *discordgo.Channel,
	user *discordgo.User, message *discordgo.MessageCreate, conf *Config, cmdHandler *CommandHandler, sessions *SessionManager) *Context {
	ctx := new(Context)
	ctx.Discord = discord
	ctx.Guild = guild
	ctx.TextChannel = textChannel
	ctx.User = user
	ctx.Message = message
	ctx.Conf = conf
	ctx.CmdHandler = cmdHandler
	ctx.Sessions = sessions
	return ctx
}

// Reply replies a basic message in the textchannel
func (ctx Context) Reply(content string) *discordgo.Message {
	msg, err := ctx.Discord.ChannelMessageSend(ctx.TextChannel.ID, content)
	if err != nil {
		fmt.Println("Error whilst sending message,", err)
		return nil
	}
	return msg
}

// GetVoiceChannel returns calling user's current voicechannel
func (ctx *Context) GetVoiceChannel() (*discordgo.Channel, error) {
	if ctx.VoiceChannel != nil {
		return ctx.VoiceChannel, nil
	}
	for _, state := range ctx.Guild.VoiceStates {
		if state.UserID == ctx.User.ID {
			channel, err := ctx.Discord.State.Channel(state.ChannelID)
			ctx.VoiceChannel = channel
			return channel, err
		}
	}
	return nil, nil
}
