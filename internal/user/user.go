package user

import "simple_chat_server/internal/message"

func (u *User) GetUserName() string {
	return u.name
}

func (u *User) SendMessageToUser(message message.IMessage) {
	u.out <- message
}

func (u *User) GetOutChannel() chan message.IMessage {
	return u.out
}

func (u *User) GetCurrentUserGroup() string {
	return u.currentGroup
}

func (u *User) SetCurrentUserGroup(groupName string) {
	u.currentGroup = groupName
}
