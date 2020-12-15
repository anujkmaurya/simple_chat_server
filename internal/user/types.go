package user

import (
	"simple_chat_server/internal/message"
)

type (
	User struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		groups       map[string]struct{}
	}
)
