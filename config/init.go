package config

import (
	"log"

	"simple_chat_server/internal/model"
)

func Init(environment string) *Config {

	cfg := &Config{}

	ok := cfg.ReadConfig(model.ConfigPath[environment])
	if !ok {
		log.Fatal("Failed to read config file")
	}
	return cfg
}
