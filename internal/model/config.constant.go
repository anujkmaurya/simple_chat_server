package model

//contains constant for config file path location
var (
	ConfigPath = map[string]string{
		"development": "etc/simple-chat-server/development/simple-chat-server.ini",
		"test":        "../internal/mock/config/config-mock.ini",
	}
)
