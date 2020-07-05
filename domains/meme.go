package domains

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

// MemeStruct maps json file into a type
//TODO: command = []string ?
type MemeStruct struct {
	Command    string `json:"command"`
	YoutubeURL string `json:"youtubeUrl"`
	FileName   string `json:"fileName"`
	Help       string `json:"help"`
}

//LoadMemes parses json file containing memes
func LoadMemes(filename string) *[]MemeStruct {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config,", err)
		return nil
	}
	var memes []MemeStruct
	json.Unmarshal(body, &memes)
	log.Info(memes)
	return &memes
}
