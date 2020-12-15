package chatmanager

import (
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
)

//IChatManager : defines all the exported methods of Chat Manager Object
type IChatManager interface {
	GetUser(userName string) (user.IUser, error)
	GetGroup(groupName string) (group.IGroup, error)
	AddGroup(group group.IGroup)
	AddUser(user user.IUser)
	RemoveUser(userName string)
	Run()
	JoinGroup(userName string, groupName string)
	SendMessageToStream(message message.IMessage)
	HandleInput(input string, userName string, groupName string) (message.IMessage, error)
}
