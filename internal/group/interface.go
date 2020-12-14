package group

import "simple_chat_server/internal/message"

type IGroup interface {
	SetGroupName(groupName string)
	GetGroupName() string
	AddUserToGroup(userName string) bool
	RemoveUserFromGroup(userName string)
	CreateSystemMessage(text string) message.IMessage
	GetSubscribedUsers() map[string]string
	GetSubscribedUsersCount() int
}
