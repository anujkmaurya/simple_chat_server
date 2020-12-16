package chatmanager

import (
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
	"sync"
)

type (
	//ChatManager : stores all the users, groups, common message stream
	ChatManager struct {
		users     *usersMap
		groupList *groupMap
		msgStream chan message.IMessage
	}

	usersMap struct {
		mutex *sync.RWMutex
		users map[string]user.IUser
	}

	groupMap struct {
		mutex  *sync.RWMutex
		groups map[string]group.IGroup
	}
)
