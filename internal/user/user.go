package user

import "simple_chat_server/internal/message"

//GetUserName : gives user name
func (u *User) GetUserName() string {
	return u.name
}

//SendMessageToUser : sends the given message to the out channel
func (u *User) SendMessageToUser(message message.IMessage) {
	u.out <- message
}

//GetOutChannel : send the output channel
func (u *User) GetOutChannel() chan message.IMessage {
	return u.out
}

//GetCurrentUserGroup: get the user's current group name
func (u *User) GetCurrentUserGroup() string {
	return u.currentGroup
}

//SetCurrentUserGroup : set the given group as the current group
//if group is new, add it to the user's subscribed group list
func (u *User) SetCurrentUserGroup(groupName string) {
	u.currentGroup = groupName

	//if group already not present, add it to users's group map
	if _, ok := u.groups[groupName]; !ok {
		u.groups[groupName] = struct{}{}
	}
}

//GetAllUserGroups : returns a map of all the subscribed groups
func (u *User) GetAllUserGroups() map[string]struct{} {
	return u.groups
}
