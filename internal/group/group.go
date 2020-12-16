package group

import (
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
)

//SetGroupName set group name
func (g *Group) SetGroupName(groupName string) {
	g.groupName = groupName
}

//GetGroupName get group name
func (g *Group) GetGroupName() string {
	return g.groupName
}

//AddUserToGroup : add userName to group
func (g *Group) AddUserToGroup(userName string) bool {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.users[userName]; !ok {
		g.users[userName] = struct{}{}
		return true
	}
	return false
}

//RemoveUserFromGroup : remove user from group
func (g *Group) RemoveUserFromGroup(userName string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if _, ok := g.users[userName]; ok {
		delete(g.users, userName)
	}
}

//CreateSystemMessage : create a system generated message to be sent to this group
func (g *Group) CreateSystemMessage(text string) message.IMessage {
	return message.CreateMessage(model.System, g.groupName, text, "")
}

//GetSubscribedUsers : return subscribed user map
func (g *Group) GetSubscribedUsers() map[string]struct{} {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	userMap := make(map[string]struct{}, len(g.users))
	for k, v := range g.users {
		userMap[k] = v
	}
	return userMap
}

//GetSubscribedUsersCount : returns counts of user subscribed to the group
func (g *Group) GetSubscribedUsersCount() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return len(g.users)
}
