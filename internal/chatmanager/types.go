package chatmanager

import (
	"errors"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
)

type (
	ChatManager struct {
		users     map[string]user.IUser
		groupList map[string]group.IGroup
		msgStream chan message.IMessage
	}
)

func (cm *ChatManager) GetUser(userName string) (user.IUser, error) {
	if val, ok := cm.users[userName]; ok {
		return val, nil
	}
	return nil, errors.New("user not present")
}

func (cm *ChatManager) GetGroup(groupName string) (group.IGroup, error) {
	if val, ok := cm.groupList[groupName]; ok {
		return val, nil
	}
	return nil, errors.New("group not present")
}

func (cm *ChatManager) AddGroup(group group.IGroup) {
	if _, ok := cm.groupList[group.GetGroupName()]; !ok {
		cm.groupList[group.GetGroupName()] = group
	}
}

func (cm *ChatManager) RemoveUser(userName string) {
	if _, ok := cm.users[userName]; ok {
		delete(cm.users, userName)
	}
	return
}

func (cm *ChatManager) SendMessageToStream(message message.IMessage) {
	cm.msgStream <- message
}

func (cm *ChatManager) AddUser(user user.IUser) {
	if _, ok := cm.users[user.GetUserName()]; !ok {
		cm.users[user.GetUserName()] = user
	}
}
