package config

import (
	"log"

	"simple_chat_server/internal/model"
)

//Init : Read configurations from config file
func Init(environment string) *Config {

	cfg := &Config{}

	ok := cfg.ReadConfig(model.ConfigPath[environment])
	if !ok {
		log.Fatal("[Err] Failed to read config file, exiting..")
	}
	return cfg
}
