package chatmanager

import (
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
)

type (
	//ChatManager: stores all the users, groups, common message stream
	ChatManager struct {
		users     map[string]user.IUser
		groupList map[string]group.IGroup
		msgStream chan message.IMessage
	}
)
