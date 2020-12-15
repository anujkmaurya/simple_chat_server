package user

import (
	"simple_chat_server/internal/message"
)

//User : user structure: stores the current group, list of subscribed groups
//out channel for sending response to the user
type (
	User struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		groups       map[string]struct{}
	}
)
