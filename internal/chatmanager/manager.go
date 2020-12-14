package chatmanager

import (
	"fmt"
	"log"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"strings"
)

func (chatManager *ChatManager) makeChannel(channelName string) {
	if _, ok := chatManager.groupList[channelName]; !ok {
		chatManager.groupList[channelName] = group.New(channelName)

		chatManager.msgStream <- chatManager.groupList["COMMON"].CreateSystemMessage(fmt.Sprintf("New Channel: %s is ready for use.", channelName))
	} else {
		chatManager.msgStream <- chatManager.groupList["COMMON"].CreateSystemMessage(fmt.Sprintf("Channel: %s already exists.", channelName))
	}
}

func (chatManager *ChatManager) JoinGroup(userName string, channelName string) {
	if _, ok := chatManager.groupList[channelName]; !ok {
		chatManager.makeChannel(channelName)
	}

	if !chatManager.groupList[channelName].AddUserToGroup(userName) {
		//log error since username is dublicate, return error
		return
	}

	chatManager.users[userName].SetCurrentUserGroup(channelName)
	chatManager.msgStream <- chatManager.groupList[channelName].CreateSystemMessage(fmt.Sprintf("%s has joint the channel. Say hello.", userName))
}

func (chatManager *ChatManager) LeaveGroup(userName string, channelName string) {
	if channelName != "COMMON" {
		chatManager.groupList[channelName].RemoveUserFromGroup(userName)

		if chatManager.users[userName].GetCurrentUserGroup() == channelName {
			user := chatManager.users[userName]
			user.SetCurrentUserGroup("COMMON")
		}

		if chatManager.groupList[channelName].GetSubscribedUsersCount() == 0 {
			//remove the group from list
			delete(chatManager.groupList, channelName)
		}

	}

}

func (chatManager *ChatManager) HandleInput(input string, userName string, channelName string) message.IMessage {
	commandArr := strings.Fields(strings.TrimSpace(input))

	if len(commandArr) == 0 {
		return message.CreateMessage(
			"System",
			"COMMON",
			"Please enter a valid string, It's empty",
			userName)
	}

	switch {
	case commandArr[0] == "--help":
		return message.CreateMessage(
			"System",
			"COMMON",
			"You can join a group using --joingroup <group name>, leave a group --leavegroup <group name> or personal to any user --personal <username>",
			userName)

	case commandArr[0] == "--personal":
		return message.CreateMessage(
			userName,
			"COMMON",
			strings.Join(commandArr[2:], " "),
			commandArr[1])

	case commandArr[0] == "--joingroup":
		chatManager.JoinGroup(userName, commandArr[1])
		return message.CreateMessage(
			"SYSTEM",
			"COMMON",
			"You successfully joined a group "+commandArr[1],
			userName)

	case commandArr[0] == "--leavegroup":
		chatManager.LeaveGroup(userName, commandArr[1])
		return message.CreateMessage(
			"SYSTEM",
			"COMMON",
			"You successfully left the group "+commandArr[1],
			userName)

	default:
		return message.CreateMessage(userName, channelName, input, "")
	}
}

func (chatManager *ChatManager) Run() {

	for {
		for message := range chatManager.msgStream {
			log.Println(message.String())
			for user := range chatManager.groupList[message.GetChannelName()].GetSubscribedUsers() {
				if _, ok := chatManager.users[user]; (message.GetReceiverName() == "" || message.GetReceiverName() == user) && ok {
					chatManager.users[user].SendMessageToUser(message)
				}
			}
		}
	}
}
