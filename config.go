package main

import (
	"encoding/json"
	"os"
)

type configFile struct {
	SteamAPIKey string
	SlackAPIURL string
	Accounts    []*configAccount
}

type configAccount struct {
	DotaID      uint64
	SlackName   string
	LastMatchID uint64
}

var config = new(configFile)

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		panic(err)
	}
}

func saveConfig() {
	file, err := os.Create("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	file.Write(data)
}
