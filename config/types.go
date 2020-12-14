package config

import (
	"time"
)

type (
	Config struct {
		Server  ServerConfig
		Timeout time.Duration
		Log     LogConfig
	}

	ServerConfig struct {
		Host              string
		AppHost           string
		Port              string
		TemplatePath      string
		Timeout           time.Duration
		PingClientTimeout time.Duration
		LocalIP           string
	}

	LogConfig struct {
		Path string
	}
)
