package chatmanager

import (
	"errors"
	"log"
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
	if user, ok := cm.users[userName]; ok {

		//find all groups of users, disconnect from all groups
		for groupName := range user.GetAllUserGroups() {

			//check if group present
			if group, err := cm.GetGroup(groupName); err == nil {

				group.RemoveUserFromGroup(userName)
			}
		}

		//delete from usersmap
		delete(cm.users, userName)
	}
	return
}

func (cm *ChatManager) SendMessageToStream(message message.IMessage) {
	cm.msgStream <- message
}

func (cm *ChatManager) AddUser(user user.IUser) {
	userName := user.GetUserName()
	if _, ok := cm.users[userName]; !ok {
		cm.users[userName] = user

		//add user to group
		//find all groups of users, disconnect from all groups
		for groupName := range user.GetAllUserGroups() {

			//check if group present
			if group, err := cm.GetGroup(groupName); err == nil {

				if !group.AddUserToGroup(userName) {
					log.Printf("[Err] Failed to add user: %s to group :%s\n", userName, groupName)
				}
			}
		}
	}
}
