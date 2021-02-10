package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

// GmoConfig has GMO Coin API connect info.
type GmoConfigList struct {
	ApiKey    string
	ApiSecret string
	BaseURL   string
}

var GmoConfig GmoConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	GmoConfig = GmoConfigList{
		ApiKey:    cfg.Section("gmo").Key("api_key").String(),
		ApiSecret: cfg.Section("gmo").Key("api_secret").String(),
		BaseURL:   cfg.Section("gmo").Key("base_url").String(),
	}
}
