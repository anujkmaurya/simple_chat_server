package model

const (
	//MaxUserMessageQueueLen :Maximum message buffer per user
	MaxUserMessageQueueLen int = 5

	//MaxChatManagerMessageQueueLen : Maximum message buffer for chat manager
	MaxChatManagerMessageQueueLen int = 10
)

const (
	//EnvDevelopemnt : dev env
	EnvDevelopemnt = "development"
)

const (
	//CommonGroup : common group name
	CommonGroup string = "COMMON"

	//System Message sender
	System = "SYSTEM"
)
