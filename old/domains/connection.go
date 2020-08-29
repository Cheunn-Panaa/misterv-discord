package domains

import (
	"github.com/bwmarrin/discordgo"
	"sync"
)

type Connection struct {
	VoiceConnection *discordgo.VoiceConnection
	send            chan []int16
	lock            sync.Mutex
	sendpcm         bool
	stopRunning     bool
	playing         bool
}

func NewConnection(voiceConnection *discordgo.VoiceConnection) *Connection {
	connection := new(Connection)
	connection.VoiceConnection = voiceConnection
	return connection
}
func (connection Connection) Disconnect() {
	connection.VoiceConnection.Disconnect()
}
