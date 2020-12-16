package user

import (
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
	"sync"
)

//New : Create new User ; assign common group to the user
func New(name string) IUser {

	user := &User{
		name:         name,
		out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
		currentGroup: model.CommonGroup,
		groups:       make(map[string]struct{}, 0),
		mutex:        &sync.RWMutex{},
	}
	user.groups[user.currentGroup] = struct{}{}
	return user
}
