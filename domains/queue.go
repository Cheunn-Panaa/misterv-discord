package domains

type SongQueue struct {
	list    []Sound
	current *Sound
	Running bool
}
