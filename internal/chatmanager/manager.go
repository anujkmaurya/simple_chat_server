package chatmanager

//this file contains all teh basic getter setters or Chatmanager struct
import (
	"errors"
	"log"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
)

//GetUser :find if a user is present with given userName
func (cm *ChatManager) GetUser(userName string) (user.IUser, error) {

	if val, ok := cm.users.getUser(userName); ok {
		return val, nil
	}
	return nil, errors.New("user not present")
}

//GetGroup :find if a group is present with given groupName
func (cm *ChatManager) GetGroup(groupName string) (group.IGroup, error) {
	if val, ok := cm.groupList.getGroup(groupName); ok {
		return val, nil
	}
	return nil, errors.New("group not present")
}

//AddGroup :Add the new group to the Group Map
func (cm *ChatManager) AddGroup(group group.IGroup) {
	// check if group is already present
	if _, ok := cm.groupList.getGroup(group.GetGroupName()); !ok {
		//if not, add it to group map
		cm.groupList.setGroup(group.GetGroupName(), group)
	}
}

//RemoveUser : removed users from users map, unsubscribes user from all the groups user was subscribed to
func (cm *ChatManager) RemoveUser(userName string) {
	if user, ok := cm.users.getUser(userName); ok {

		//find all groups of users, disconnect from all groups
		for groupName := range user.GetAllUserGroups() {

			//check if group present
			if group, err := cm.GetGroup(groupName); err == nil {

				group.RemoveUserFromGroup(userName)
			}
		}

		//delete user from usersmap
		cm.users.deleteUser(userName)
	}
	return
}

//SendMessageToStream : send message to Message Stream
func (cm *ChatManager) SendMessageToStream(message message.IMessage) {
	cm.msgStream <- message
}

//AddUser : Add user to the group
func (cm *ChatManager) AddUser(user user.IUser) {
	userName := user.GetUserName()
	if _, ok := cm.users.getUser(userName); !ok {
		cm.users.setUser(userName, user)

		//find all groups of users, add user to all groups he wants to subscribe to
		for groupName := range user.GetAllUserGroups() {

			//check if group present
			if group, err := cm.GetGroup(groupName); err == nil {

				//add user to group
				if !group.AddUserToGroup(userName) {
					log.Printf("[Err] Failed to add user: %s to group :%s\n", userName, groupName)
				}
			}
		}
	}
}
