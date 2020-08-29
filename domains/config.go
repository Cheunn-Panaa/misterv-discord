package domains

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)
type BotConf struct {
	Token	string `yaml:"token"`
	Name	string `yaml:"name"`
	Prefix	string `yaml:"prefix"`
	Playing string `yaml:"playing"`
	Status 	string `yaml:"status"`
	Help    string `yaml:"help"`
}
// Config is configuration of Bot
type Config struct {
	Bot BotConf
	Discord struct {
		ApplicationId	int `yaml:"appliationID"`
		LogChannelId	int `yaml:"logChannelId"`
		Sharding struct {
			Id	int `yaml:"id"`
		}
	}
	Sound struct {
		SoundDir       string `yaml:"directory"`
		SoundCacheSize int    `yaml:"soundCacheSize"`
		MaxQueueSize   int    `yaml:"maxQueueSize"`
	}
	Env            string `yaml:"env"`
}

// NewConfig is constructor
func NewConfig(filename string) (config *Config, err error) {
	b, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(b, &config)
	return
}