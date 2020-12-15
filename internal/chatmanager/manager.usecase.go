package chatmanager

import (
	"errors"
	"fmt"
	"log"
	"simple_chat_server/internal/command"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
	"strings"
)

//createNewGroup :creates a new group on receiving create a new group command
func (chatManager *ChatManager) createNewGroup(userName, groupName string) bool {
	if _, ok := chatManager.groupList[groupName]; !ok {
		//create a new group instance
		chatManager.groupList[groupName] = group.New(groupName)
		//send new group created message to all members of this group
		chatManager.SendMessageToStream(chatManager.groupList[model.CommonGroup].CreateSystemMessage(fmt.Sprintf("New Group: %s is ready for use.", groupName)))

		// chatManager.msgStream <- chatManager.groupList[model.CommonGroup].CreateSystemMessage(fmt.Sprintf("New Group: %s is ready for use.", groupName))
		return true
	}
	//send response to user to choose a new group name
	chatManager.SendMessageToStream(chatManager.groupList[model.CommonGroup].CreateSystemMessage(fmt.Sprintf("Group: %s already exists. Please choose another name", groupName)).SetReceiverName(userName))

	// chatManager.msgStream <- chatManager.groupList[model.CommonGroup].CreateSystemMessage(fmt.Sprintf("Group: %s already exists. Please choose another name", groupName))

	return false
}

//JoinGroup : allows a user to join a group, creates a new group if it doesn't exist,
// and assigns given group as current group
func (chatManager *ChatManager) JoinGroup(userName string, groupName string) {

	//if new group, create new group
	if _, ok := chatManager.groupList[groupName]; !ok {
		if !chatManager.createNewGroup(userName, groupName) {
			return
		}
	}

	//add user to the grouplist
	chatManager.groupList[groupName].AddUserToGroup(userName)

	//make given group as current group, switch user to this group
	chatManager.users[userName].SetCurrentUserGroup(groupName)

	//broadcast new user joining message to all users in this group
	chatManager.SendMessageToStream(chatManager.groupList[groupName].CreateSystemMessage(fmt.Sprintf("%s has joint the channel. Say hello.", userName)))
}

//LeaveGroup : UnSubscribe user from a given group
func (chatManager *ChatManager) LeaveGroup(userName string, groupName string) {
	//user can't unsubscribe from common group
	if groupName != model.CommonGroup {
		//remove user from group
		chatManager.groupList[groupName].RemoveUserFromGroup(userName)

		//if user's current group was given group
		if chatManager.users[userName].GetCurrentUserGroup() == groupName {

			user := chatManager.users[userName]
			//make common channel as the current group
			user.SetCurrentUserGroup(model.CommonGroup)
		}

		//if group doesn't have any user, delete group
		if chatManager.groupList[groupName].GetSubscribedUsersCount() == 0 {
			//remove the group from group map
			delete(chatManager.groupList, groupName)
		}

	}

}

//HandleInput : handles the user's input from terminal
//acts as a command centre: parses commands, takes action on th basis of that
func (chatManager *ChatManager) HandleInput(input string, userName string, groupName string) (message.IMessage, error) {

	commandArr := strings.Fields(strings.TrimSpace(input))
	//check for empty input from user
	if len(commandArr) == 0 {
		//inform user and ask to enter correct input
		return message.CreateMessage(
			model.System,
			model.CommonGroup,
			"Please enter a valid string, It's empty",
			userName), nil
	}

	//parse and sanitise command received
	if command.ParseAndSanitiseCommand(commandArr) {

		userCommand := command.ParseCommand(commandArr)

		//take actions based on command received
		switch userCommand {
		case command.HelpCommand:
			return message.CreateMessage(
				model.System,
				model.CommonGroup,
				"You can join a group using --joingroup <group name>, leave a group --leavegroup <group name>, ignore an user --ignoreuser <user name>, unignore an user --unignoreuser <user name> or personal to any user --personal <username>",
				userName), nil

		case command.PersonalCommand:
			if _, err := chatManager.GetUser(commandArr[1]); err != nil {
				return nil, errors.New("receipient user is not present")
			} else {
				return message.CreateMessage(
					userName,
					model.CommonGroup,
					strings.Join(commandArr[2:], " "),
					commandArr[1]), nil

			}

		case command.JoinGroupCommand:
			chatManager.JoinGroup(userName, commandArr[1])

			return message.CreateMessage(
				model.System,
				model.CommonGroup,
				"You successfully joined a group "+commandArr[1],
				userName), nil

		case command.LeaveGroupCommand:

			chatManager.LeaveGroup(userName, commandArr[1])
			return message.CreateMessage(
				model.System,
				model.CommonGroup,
				"You successfully left the group "+commandArr[1],
				userName), nil

		case command.NormalMessage:
			return message.CreateMessage(userName, groupName, input, ""), nil

		case command.IgnoreUserCommand:

			chatManager.ignoreUser(userName, commandArr[1])
			return message.CreateMessage(
				model.System,
				model.CommonGroup,
				"You have successfully ignored the user "+commandArr[1],
				userName), nil

		case command.UnIgnoreUserCommand:

			chatManager.unIgnoreUser(userName, commandArr[1])
			return message.CreateMessage(
				model.System,
				model.CommonGroup,
				"You have successfully unignored the user "+commandArr[1],
				userName), nil
		}
	}

	//incorrect command
	return nil, errors.New("command sanitisation failed")
}

//Run : reads continuously from the message streams and relay it to the users
func (chatManager *ChatManager) Run() {

	for {
		//for each message through the msgStream channel
		for message := range chatManager.msgStream {
			//print message in log
			log.Print(message.String())

			//get receiver and sender name
			receiverName := message.GetReceiverName()
			senderName := message.GetSenderName()

			//loop for all the users subscribed to the group
			for user := range chatManager.groupList[message.GetGroupName()].GetSubscribedUsers() {
				//check for avoiding sending his message to himself

				if user != senderName {
					//check if the user exists in the chat mangager users map,
					//check if the message is general, or for this user
					if recipient, ok := chatManager.users[user]; (receiverName == "" || receiverName == user) && ok {

						//check for blocked/ignored user
						fmt.Println(recipient.GetIgnoredUserName())
						if recipient.GetIgnoredUserName() != senderName {
							recipient.SendMessageToUser(message)
						}
					}
				}
			}
		}
	}
}

//IgnoreUser : allows a user to stop receiving message from a sender other than system
func (chatManager *ChatManager) ignoreUser(userName string, ignoredUserName string) {

	if user, ok := chatManager.users[userName]; ok {
		if _, isPresent := chatManager.users[ignoredUserName]; isPresent {
			user.SetIgnoredUserName(ignoredUserName)
		}
	}
}

//unIgnoreUser : allow user from to receive message
func (chatManager *ChatManager) unIgnoreUser(userName string, ignoredUserName string) {

	if user, ok := chatManager.users[userName]; ok {
		if user.GetIgnoredUserName() == ignoredUserName {
			user.SetIgnoredUserName("")
		}
	}
}
