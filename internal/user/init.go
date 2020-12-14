package user

import (
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
)

func New(name string) IUser {

	return &User{
		name:         name,
		out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
		currentGroup: "COMMON",
	}

}
