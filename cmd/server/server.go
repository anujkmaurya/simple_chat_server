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
	log.Println("starting server")

	//read Environment variable :APPENV
	environment := os.Getenv("APPENV")
	if environment == "" {
		environment = model.EnvDevelopemnt
	}

	//init configurations, read from config file
	cfg := config.Init(environment)

	//init logger instance
	err := logger.InitLogger(cfg.Log.Path)
	if err != nil {
		log.Fatalf("Failed to initialise logger %+v\n", err)
	}

	//server port
	port := cfg.Server.Port

	//server starts listening the configured port
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln("[Err] Server Failed to Listen, err:", err.Error())
	}
	defer server.Close()

	//init chatMangager
	chatManager := chatmanager.New()

	//add a default group "Common"
	chatManager.AddGroup(group.New(model.CommonGroup))

	//spawn Run function in a goroutine
	go chatManager.Run()

	//loop to accept new connections
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln("[Err] Server Failed to Accept new connection, err:", err.Error())
		}

		//handle a succesfull user connection in a goroutine
		go HandleUserConnection(chatManager, conn)
	}
}

//HandleUserConnection : handles user connection,
func HandleUserConnection(chatManager chatmanager.IChatManager, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var userName string
	var iGroup group.IGroup
	var err error
	if iGroup, err = chatManager.GetGroup(model.CommonGroup); err != nil {
		log.Printf("[Err] error in getting groupName: %s, err:%+v\n", model.CommonGroup, err)
		return
	}

	//Welcome message
	if _, err := io.WriteString(conn, iGroup.CreateSystemMessage("Welcome to Chat Groups, please choose an username: ").String()); err != nil {
		log.Println("[Err] error in writing to the connection")
	}
	//read Conn for userName
	for {
		if scanner.Scan() {
			//get username
			userName = scanner.Text()
			if _, err := chatManager.GetUser(userName); err != nil {
				//create new user and add to chatmanager
				chatManager.AddUser(user.New(userName))

				if _, err := io.WriteString(conn, iGroup.CreateSystemMessage(fmt.Sprintf("Thanks %s! for joining us. Type --help for a list of commands.", userName)).String()); err != nil {
					log.Println("[Err] error in writing to the connection")
				}

				break
			}
			//ask user to choose another name in case
			if _, err := io.WriteString(conn, iGroup.CreateSystemMessage(fmt.Sprintf("Sorry the user name: %s is taken. Please choose another one", userName)).String()); err != nil {
				log.Println("[Err] error in writing to the connection")
			}

		} else {
			//handles the case if client disconnects while/before adding username
			log.Println("[Err] A new client disconnected while adding username")
			return
		}
	}

	//by default, user joins the Common group
	chatManager.JoinGroup(userName, model.CommonGroup)

	//scan all inputs from this user in a go routine
	go func() {
		//remove user, if TCP connection breaks or client closes the connection
		defer func() {
			log.Println("[Err] client disconnected, hence removing user: ", userName)
			chatManager.RemoveUser(userName)
		}()

		//scan for user message continuously and send Message to chatmanager
		for scanner.Scan() {
			//check for scanner error
			if scanner.Err() != nil {
				return
			}

			input := scanner.Text()
			if user, err := chatManager.GetUser(userName); err == nil {
				if msg, err := chatManager.HandleInput(input, userName, user.GetCurrentUserGroup()); err == nil {

					//send message to message stream
					chatManager.SendMessageToStream(msg)
				} else {

					//ask this user to renter the command or new text
					if _, err := io.WriteString(conn, iGroup.CreateSystemMessage("The command is incorrect,  Type --help for a list of commands.").SetReceiverName(userName).String()); err != nil {
						log.Println("[Err] error in writing to the connection")
					}
				}
			}
		}
	}()

	//for all the messages, to be transmitted to this user
	if user, err := chatManager.GetUser(userName); err == nil {
		for message := range user.GetOutChannel() {
			if _, err := io.WriteString(conn, message.String()); err != nil {
				log.Println("[Err] error in writing to the connection")
			}
		}
	} else {
		log.Printf("[Err] user: %s is not present inn the system ", userName)
	}
}
