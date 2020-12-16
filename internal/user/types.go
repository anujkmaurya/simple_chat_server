package user

import (
	"simple_chat_server/internal/message"
	"sync"
)

//User : user structure: stores the current group, list of subscribed groups
//out channel for sending response to the user
type (
	User struct {
		mutex        *sync.RWMutex
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
	}
)
