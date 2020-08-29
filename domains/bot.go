package domains

import (
	log "github.com/Sirupsen/logrus"
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)
var once sync.Once
// Bot is instance of bot
type Bot struct {
	Session *discordgo.Session
	config BotConf
	playerInstance map[string]*Player
	//queues map[string][]*Play
	//m      *sync.Mutex
	//cache  *cache.Cache
}
var (
	config     *Config
	bot	*Bot
	mutex             sync.Mutex
	//botID      string
)

// NewBot is constructor
func NewBot(config *Config) (b *Bot ,err error) {
	discord, err := discordgo.New(config.Bot.Token)

	once.Do(func() { // <-- atomic, does not allow repeating
		bot = &Bot{
			Session: discord,
			config: config.Bot,
			playerInstance: map[string]*Player{},
		} // <-- thread safe
	})
	//mutex.Lock()
	//bot = &Bot{
	//	Session:        discord,
	//	config:         config.Bot,
	//	playerInstance: map[string]*Player{},
	//} // <-- thread safe
	//mutex.Unlock()
	bot.Session.AddHandler(bot.ready)
	bot.Session.AddHandler(bot.commandHandler)

	//bot.AddHandler(bot.messageCreate)

	//discord.cache.OnEvicted(func(key string, data interface{}) {
	//	log.Printf("Evicted cache:%s\n", key)
	//})
	return bot, bot.Session.Open()
}



func (b *Bot) ready(s *discordgo.Session, event *discordgo.Ready) {
	_ = s.UpdateStatus(0, b.config.Status)
}

func (b *Bot) commandHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !strings.HasPrefix(message.Content, b.config.Prefix) {
		return
	} else {
		log.Info("message received: " + message.Content)

		player := b.playerInstance[message.GuildID]
		b.joinChannel(player, message, session)
		_ , err := b.PlaySound(b.playerInstance[message.GuildID].voice)
		if err != nil {
			return
		}
		//joinChannel(session,message)
		/*guild, _ := session.State.Guild(message.GuildID)
		vc, _ := GetVoiceChannel(session, guild, message.Author)

		_, _ = session.ChannelMessageSend(message.ChannelID, b.config.Help)*/
		return
	}

}
func GetVoiceChannel(user string) (voiceChannel string) {

	for _, g := range bot.Session.State.Guilds {
		for _, v := range g.VoiceStates {
			if v.UserID == user {
				return v.ChannelID
			}
		}
	}
	return ""
}
func (b *Bot) joinChannel(player *Player, message *discordgo.MessageCreate, session *discordgo.Session){
	log.Info("Joining Channel request from ", message.Author.Username)

	if player != nil {
		log.Debug("INFO: Player Instance already created.")
	} else {
		guildID := b.GetGuild(message.ChannelID)
		// create new voice instance
		mutex.Lock()
		player = new(Player) // <-- thread safe
		player.guildID = guildID
		player.session = session
		b.playerInstance[guildID] = player
		log.Info(b.playerInstance)
		mutex.Unlock()
	}
		log.Info("Not in a voice channel, joining ...")
		voiceChannelID := GetVoiceChannel( message.Author.ID)
		if voiceChannelID == "" {
			log.Info("Can't join voice channel, user out of the discord")
			return
		}

		var err error
		log.Info(b.Session, player.session)
		player.voice, err = b.Session.ChannelVoiceJoin(player.guildID, voiceChannelID, false, true)
		if err != nil {
			//player.Stop()
			log.Println("ERROR: Error to join in a voice channel: ", err)
			return
		}
		_ = player.voice.Speaking(false)
		log.Debug("INFO: New Voice Instance created")

	//TODO check if possible
	//voiceChannelId := GetVoiceChannel(message.Author.ID);
	//if voiceChannelId == "" {
	//	log.Error("ERROR: Voice channel id not found.")
	//	// Send response
	//	//ChMessageSend(m.ChannelID, "[**Music**] <@"+ m.Author.ID +"> You need to join a voice channel!")
	//	return
	//}


}
func (b *Bot) PlaySound(vc *discordgo.VoiceConnection) (*discordgo.VoiceConnection, error) {
	var data [][]byte

	err := vc.Speaking(true)
	if err != nil {

		log.Info("err speak" )
		return nil, err
	}
	// Load from file
	f, err := os.Open(filepath.Join("/opt/app", "tuveuxquoi.dca"))
	if err != nil {
		log.Info("err join" , err)
		return nil, err
	}
	defer f.Close()


	//if err = b.sendSilence(vc, 10); err != nil {
	//	return nil, err
	//}
	decoder := dca.NewDecoder(f)
	for {
		frame, err := decoder.OpusFrame()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		data = append(data, frame)
		vc.OpusSend <- frame
	}

	//if err = b.sendSilence(vc, 5); err != nil {
	//	return nil, err
	//}
	//
	//if b.cache.ItemCount() < b.config.SoundCacheSize {
	//	b.cache.SetDefault(play.Sound.Name, data)
	//}

	return vc, err
}

// SearchGetGuild search the guild ID
func (b *Bot) GetGuild(textChannelID string) (guildID string) {
	channel, _ := b.Session.Channel(textChannelID)
	guildID = channel.GuildID
	return
}
//func registerAllCommands() {
//	log.Info("Registering all commands")
//	cmdHandler.Register("meme", commands.MemeCommand, "LA FETE")
//
//	log.Debug("Loading Memes")
//	memes := domains.LoadMemes(os.Getenv("MEME_FILE"))
//	log.WithFields(log.Fields{
//		"memes": memes,
//	}).Debug("Memes loaded")
//	for index, meme := range *memes {
//		cmdHandler.RegisterMemeCmd(meme.Command, commands.MemeCommand, meme.Help, meme.YoutubeURL, meme.FileName)
//		log.WithFields(log.Fields{
//			"index": index,
//			"memes": meme.Command,
//		}).Debug("Registering meme command")
//	}
//}
