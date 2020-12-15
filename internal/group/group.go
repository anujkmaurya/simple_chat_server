package group

import (
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
)

func (g *Group) SetGroupName(groupName string) {
	g.groupName = groupName
}

func (g *Group) GetGroupName() string {
	return g.groupName
}

func (g *Group) AddUserToGroup(userName string) bool {
	if _, ok := g.users[userName]; !ok {
		g.users[userName] = struct{}{}
		return true
	}
	return false
}

func (g *Group) RemoveUserFromGroup(userName string) {
	if _, ok := g.users[userName]; ok {
		delete(g.users, userName)
	}
}

func (group *Group) CreateSystemMessage(text string) message.IMessage {
	return message.CreateMessage(model.System, group.groupName, text, "")
}

func (g *Group) GetSubscribedUsers() map[string]struct{} {
	return g.users
}

func (g *Group) GetSubscribedUsersCount() int {
	return len(g.users)
}
