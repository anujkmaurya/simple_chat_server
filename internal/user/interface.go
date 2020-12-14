package user

import "simple_chat_server/internal/message"

type IUser interface {
	GetUserName() string
	SendMessageToUser(message message.IMessage)
	GetOutChannel() chan message.IMessage
	GetFocusedGroup() string
}
