package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Prefix        string `json:"bot_prefix"`
	BotName       string `json:"bot_name"`
	BotToken      string `json:"bot_token"`
	OwnerId       string `json:"owner_id"`
	UseSharding   bool   `json:"use_sharding"`
	ShardId       int    `json:"shard_id"`
	ShardCount    int    `json:"shard_count"`
	DefaultStatus string `json:"default_status"`
}

func LoadConfig(filename string) *Config {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error loading config,", err)
		return nil
	}
	var conf Config
	json.Unmarshal(body, &conf)
	return &conf
}
