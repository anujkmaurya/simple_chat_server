package config

type (
	Config struct {
		Server ServerConfig
		Log    LogConfig
	}

	ServerConfig struct {
		Host string
		Port string
	}

	LogConfig struct {
		Path string
	}
)
