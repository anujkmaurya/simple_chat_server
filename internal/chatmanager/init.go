package chatmanager

import (
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
	"simple_chat_server/internal/user"
	"sync"
)

//New : Create new Chat Manager
func New() IChatManager {

	return &ChatManager{
		users: &usersMap{
			mutex: &sync.RWMutex{},
			users: make(map[string]user.IUser, 0),
		},
		groupList: &groupMap{
			mutex:  &sync.RWMutex{},
			groups: make(map[string]group.IGroup, 0),
		},
		msgStream: make(chan message.IMessage, model.MaxChatManagerMessageQueueLen),
	}
}
