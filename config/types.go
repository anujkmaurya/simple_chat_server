package config

import "time"

type (
	//Config : stores all configurations
	Config struct {
		Server ServerConfig
		Log    LogConfig
		Client ClientConfig
	}

	//ServerConfig : all server configs
	ServerConfig struct {
		Host string
		Port string
	}

	//ClientConfig : all client configs
	ClientConfig struct {
		TimeOut time.Duration
	}

	//LogConfig : contains logger specific configs
	LogConfig struct {
		Path string
	}
)
