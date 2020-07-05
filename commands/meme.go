package commands

import (
	"../domains"
	"bytes"
	log "github.com/Sirupsen/logrus"
	"os/exec"
)

// MemeCommand Handles majority of the work for meme commands defined in the json file
func MemeCommand(ctx domains.Context) {

	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	joinChannel(ctx)

	cmd := exec.Command("youtube-dl --audio-format opus -o - D530X1eRJAk")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {

		log.Fatal(err)
	}
	log.Info(out.String())
	log.Info(out)
	sess.Connection.VoiceConnection.OpusSend <- out.Bytes()
	//cmd := exec.Command("youtube-dl", "--skip-download", "--print-json", )
}

func joinChannel(ctx domains.Context) {
	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)

	if sess == nil {
		log.Info("Not in a voice channel, joining ...")
		vc, _ := ctx.GetVoiceChannel()
		if vc == nil {
			log.Info("Can't join voice channel, user out of the discord")
			ctx.Reply(ctx.Message.Author.Mention() + " Tu dois être dans un channel vocal pour m'appeler fréro")
			return
		}
		sess, err := ctx.Sessions.Join(ctx.Discord, ctx.Guild.ID, vc.ID, domains.JoinProperties{
			Muted:    false,
			Deafened: true,
		})

		if err != nil {
			log.Fatal(err)
			return
		}
		log.WithField("channelID", sess.ChannelId).Info("Joined voice channel")
	}
}
