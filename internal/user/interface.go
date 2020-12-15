package user

import "simple_chat_server/internal/message"

//IUser :defines all the exported methods of User Object
type IUser interface {
	GetUserName() string
	SendMessageToUser(message message.IMessage)
	GetOutChannel() chan message.IMessage
	GetCurrentUserGroup() string
	SetCurrentUserGroup(groupName string)
	GetAllUserGroups() map[string]struct{}
	GetIgnoredUserName() string
	SetIgnoredUserName(userName string)
}
