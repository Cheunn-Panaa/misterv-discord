package commands

import (
	"../domains"
	"encoding/binary"
	log "github.com/Sirupsen/logrus"
	"io"
	"os"
)

var itagValues = [18]int{249, 250, 251, 13, 17, 18, 22, 34, 35, 36, 37, 38, 83, 84, 85, 139, 140, 141}

var buffer = make([][]byte, 0)
// MemeCommand Handles majority of the work for meme commands defined in the json file
func MemeCommand(ctx domains.Context) {
//	sess := ctx.Sessions.GetByGuild(ctx.Guild.ID)
	//joinChannel(ctx)
	vc, _ := ctx.GetVoiceChannel()
	sess, err := ctx.Discord.ChannelVoiceJoin(ctx.Guild.ID, vc.ID ,false,true)
	if err != nil {
		return
	}

	// Start speaking
	_ = sess.Speaking(true)


	// Stop speaking
	_ = sess.Speaking(false)

	file, err := os.Open("../src/tuveuxquoitoi.dca")
	if err != nil {
		return
	}

	var opuslen int16
	// Send the buffer data.
	for _, buff := range buffer {
		sess.OpusSend <- buff
	}
	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return
			}
			return
		}

		if err != nil {
			return
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			return
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)
	}


	//TODO: https://github.com/kkdai/youtube
	//cmd := exec.Command("./usr/local/bin/youtube-dl", "--audio-format", "opus", "-o","-", "D530X1eRJAk")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//test := cmd.Run()
	//if test != nil {
	//	log.Fatal(test)
	//}
	//_ = cmd.Wait()
	//sess.OpusSend <- out.Bytes()


	//videoID := "D530X1eRJAk"
	//client := youtube.Client{}
	//
	//var data [][]byte
	//video, err := client.GetVideo(videoID)
	//for i, format := range video.Formats {
	//	for _, item := range itagValues {
	//		if format.ItagNo == item {
	//			resp, err := client.GetStream(video, &video.Formats[i])
	//			if err != nil {
	//				return
	//			}
	//			var url = resp.Body;
	//			defer url.Close();
	//			options := dca.StdEncodeOptions
	//			options.RawOutput = true
	//			options.Bitrate = 96
	//			options.Application = "lowdelay"
	//			dca.EncodeFile(url, options)
	//
	//			encode, _ := dca.EncodeFile(url, options)
	//			defer encode.Cleanup()
	//			for {
	//					var sz_frame int16
	//					err := binary.Read(encode, binary.LittleEndian, &sz_frame)
	//					if err != nil {
	//						return
	//					}
	//					Inbuf := make([]byte, sz_frame)
	//					_ = binary.Read(encode, binary.LittleEndian, &Inbuf)
	//					vc.OpusSend <- Inbuf
	//				}
	//			}
	//		}
	//	}


	//
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
