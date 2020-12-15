package user

import "simple_chat_server/internal/message"

type IUser interface {
	GetUserName() string
	SendMessageToUser(message message.IMessage)
	GetOutChannel() chan message.IMessage
	GetCurrentUserGroup() string
	SetCurrentUserGroup(groupName string)
	GetAllUserGroups() map[string]struct{}
}
