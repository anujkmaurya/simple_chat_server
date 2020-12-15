package chatmanager

import (
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
)

//New: Create new Chatmanager, it's unique through out program
func New() IChatManager {

	return &ChatManager{
		users:     make(map[string]user.IUser, 0),
		groupList: make(map[string]group.IGroup, 0),
		msgStream: make(chan message.IMessage, 5),
	}
}
