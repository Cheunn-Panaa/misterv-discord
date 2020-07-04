package domains

import (
	"github.com/bwmarrin/discordgo"
	_ "github.com/bwmarrin/discordgo"
)

type Play struct {
	GuildID   string
	ChannelID string
	UserID    string
	Sound     *Sound
}

func (ctx *Context) createPlay(user *discordgo.User, guild *discordgo.Guild, sound *Sound) (*Play, error) {
	channel, err := ctx.GetVoiceChannel()
	if err != nil {
		return nil, err
	}
	play := &Play{
		GuildID:   guild.ID,
		ChannelID: channel.ID,
		UserID:    user.ID,
		Sound:     sound,
	}
	return play, err
}
