package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"simple_chat_server/config"
	"simple_chat_server/internal/chatmanager"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/model"
	"simple_chat_server/internal/model/logger"
	"simple_chat_server/internal/user"

	"log"
	"net"
)

func main() {

	startApp()
}

func startApp() {
	fmt.Println("starting server")

	environment := os.Getenv("APPENV")
	if environment == "" {
		environment = model.EnvDevelopemnt
	}

	cfg := config.Init(environment)

	err := logger.InitLogger(cfg.Log.Path)
	if err != nil {
		log.Fatalf("Failed to initialise logger %+v\n", err)
	}

	port := cfg.Server.Port

	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	log.Println("server started ")

	chatManager := chatmanager.New()

	chatManager.AddGroup(group.New("COMMON"))

	go chatManager.Run()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go HandleUserConnection(chatManager, conn)
	}
}

func HandleUserConnection(chatManager chatmanager.IChatManager, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var userName string
	var iGroup group.IGroup
	var err error
	if iGroup, err = chatManager.GetGroup("COMMON"); err != nil {
		log.Printf("error in getting groupName, err:%+v\n", err)
		return
	}

	io.WriteString(conn, iGroup.CreateSystemMessage("Welcome to Chat Groups, please choose an username: ").String())

	for {
		scanner.Scan()
		userName = scanner.Text()
		fmt.Println("username: ", userName)

		if _, err := chatManager.GetUser(userName); err != nil {

			chatManager.AddUser(user.New(userName))

			io.WriteString(conn, iGroup.CreateSystemMessage("Thanks for joining us. Type /help for a list of commands.").String())

			break
		}
		io.WriteString(conn, iGroup.CreateSystemMessage("Sorry that user name is taken Please choose another one:").String())
	}

	defer func() {
		chatManager.RemoveUser(userName)
	}()

	chatManager.JoinGroup(userName, "COMMON")

	go func() {
		for scanner.Scan() {
			input := scanner.Text()
			if user, err := chatManager.GetUser(userName); err == nil {
				chatManager.SendMessageToStream(chatManager.HandleInput(input, userName, user.GetCurrentUserGroup()))
			}
		}
	}()

	if user, err := chatManager.GetUser(userName); err == nil {
		for message := range user.GetOutChannel() {
			io.WriteString(conn, message.String())
		}
	}
}
