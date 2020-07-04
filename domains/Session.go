package domains

type (
	Session struct {
		Queue              *SongQueue
		guildId, ChannelId string
		connection         *Connection
	}

	SessionManager struct {
		sessions map[string]*Session
	}

	JoinProperties struct {
		Muted    bool
		Deafened bool
	}
)
