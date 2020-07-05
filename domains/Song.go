package domains

type Song struct {
	Id       string
	Media    string
	Title    string
	Duration *string
}

func NewSong(media, title, id string) *Song {
	song := new(Song)
	song.Media = media
	song.Title = title
	song.Id = id
	return song
}
