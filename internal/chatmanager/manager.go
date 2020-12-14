package chatmanager

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/user"
	"strings"
)

func (chatManager *ChatManager) makeChannel(channelName string) {
	if _, ok := chatManager.groupList[channelName]; !ok {
		chatManager.groupList[channelName] = group.New(channelName)

		chatManager.msgStream <- chatManager.groupList["GENERAL"].CreateSystemMessage(fmt.Sprintf("New Channel: %s is ready for use.", channelName))
	} else {
		chatManager.msgStream <- chatManager.groupList["GENERAL"].CreateSystemMessage(fmt.Sprintf("Channel: %s already exists.", channelName))
	}
}

func (chatManager *ChatManager) JoinChannel(userName string, channelName string) {
	if _, ok := chatManager.groupList[channelName]; !ok {
		chatManager.makeChannel(channelName)
	}

	if !chatManager.groupList[channelName].AddUserToGroup(userName) {
		//log error since username is dublicate, return error
		return
	}

	// chatManager.users[userName].focused = channelName
	chatManager.msgStream <- chatManager.groupList[channelName].CreateSystemMessage(fmt.Sprintf("%s has joint the channel. Say hello.", userName))
}

func (chatManager *ChatManager) UnJoinChannel(userName string, channelName string) {
	if channelName != "GENERAL" {
		chatManager.groupList[channelName].RemoveUserFromGroup(userName)

	}
}

func (chatManager *ChatManager) HandleInput(input string, userName string, channelName string) message.IMessage {
	commandArr := strings.Fields(input)
	switch {
	case commandArr[0] == "/help":
		return message.CreateMessage(
			"System",
			"GENERAL",
			"You can join a channel /join <channel>, unjoin a channel /unjoin <channel> or wisper to any user /w <username>")
	// case commandArr[0] == "/w":
	// 	return msg{
	// 		text:        strings.Join(commandArr[2:], " "),
	// 		sender:      userName,
	// 		receiver:    commandArr[1],
	// 		channelName: "GENERAL",
	// 		timeStamp:   time.Now(),
	// 	}
	// case commandArr[0] == "/join":
	// 	chatManager.joinChannel(userName, commandArr[1])
	// 	return msg{
	// 		text:        "You successfully joined a " + commandArr[1],
	// 		sender:      "SYSTEM",
	// 		receiver:    userName,
	// 		channelName: "GENERAL",
	// 		timeStamp:   time.Now(),
	// 	}
	// case commandArr[0] == "/unjoin":
	// 	chatManager.unJoinChannel(userName, commandArr[1])
	// 	return msg{
	// 		text:        "You successfully unjoined the channel " + commandArr[1],
	// 		sender:      "SYSTEM",
	// 		receiver:    userName,
	// 		channelName: "GENERAL",
	// 		timeStamp:   time.Now(),
	// 	}
	default:
		return message.CreateMessage(userName, channelName, input)
	}
}

func (chatManager *ChatManager) Run() {
	// logFile, err := os.OpenFile(os.Getenv("LOG_FILE_LOCATION"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// logFile, err := os.OpenFile("abc.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer logFile.Close()
	// log.SetOutput(logFile)
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

func HandleUserConnection(chatManager IChatManager, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var userName string
	var iGroup group.IGroup
	var err error
	if iGroup, err = chatManager.GetGroup("GENERAL"); err != nil {
		log.Printf("error in getting groupName, err:%+v\n", err)
		return
	} else {

		io.WriteString(conn, iGroup.CreateSystemMessage("Welcome to GoChatWin an awesome chat server pleas chose a UserName: ").String())
	}

	// check users will be unique and not overwrite eachother.
	for {
		scanner.Scan()
		userName = scanner.Text()
		if _, err := chatManager.GetUser(userName); err == nil {

			chatManager.AddUser(user.New(userName))

			io.WriteString(conn, iGroup.CreateSystemMessage("Thanks for joining us. Type /help for a list of commands.").String())

			break
		}
		io.WriteString(conn, iGroup.CreateSystemMessage("Sorry that user name is taken Please choose another one:").String())
	}

	defer func() {
		chatManager.RemoveUser(userName)
	}()

	chatManager.JoinChannel(userName, "GENERAL")

	go func() {
		for scanner.Scan() {
			input := scanner.Text()
			if user, err := chatManager.GetUser(userName); err == nil {
				chatManager.SendMessageToStream(chatManager.HandleInput(input, userName, user.GetFocusedGroup()))
			}
		}
	}()

	if user, err := chatManager.GetUser(userName); err == nil {
		for message := range user.GetOutChannel() {
			io.WriteString(conn, message.String())
		}
	}
}
