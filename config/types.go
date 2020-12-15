package config

import "time"

type (
	Config struct {
		Server ServerConfig
		Log    LogConfig
		Client ClientConfig
	}

	ServerConfig struct {
		Host string
		Port string
	}

	ClientConfig struct {
		TimeOut time.Duration
	}

	LogConfig struct {
		Path string
	}
)
