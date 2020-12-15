package group

import "simple_chat_server/internal/message"

//IGroup : defines all the exported methods of Group Object
type IGroup interface {
	SetGroupName(groupName string)
	GetGroupName() string
	AddUserToGroup(userName string) bool
	RemoveUserFromGroup(userName string)
	CreateSystemMessage(text string) message.IMessage
	GetSubscribedUsers() map[string]struct{}
	GetSubscribedUsersCount() int
}
