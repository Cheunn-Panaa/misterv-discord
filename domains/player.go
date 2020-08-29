package domains

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"os/exec"
	"sync"
)

type Player struct {
	voice               *discordgo.VoiceConnection
	session             *discordgo.Session
	encoder             *dca.EncodeSession
	stream              *dca.StreamingSession
	run                 *exec.Cmd
	queueMutex          sync.Mutex
	audioMutex          sync.Mutex
	nowPlaying          Song
	queue               SongQueue
	recv                []int16
	guildID             string
	channelID           string
	speaking            bool
	pause               bool
	stop                bool
	skip                bool
	radioFlag           bool
}


