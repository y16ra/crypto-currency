package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type TradeConfigList struct {
	Duration string //map[string]time.Duration
}

var TradeConfig TradeConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	TradeConfig = TradeConfigList{
		Duration: cfg.Section("trading").Key("trade_duration").String(),
	}
}
