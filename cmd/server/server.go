package main

import (
	"simple_chat_server/internal/chatmanager"
	"simple_chat_server/internal/group"

	// "fmt"

	"log"
	"net"
	// "os"
	// "strings"
	// "time"
	// "simple_chat_server/internal/model/user"
	// "simple_chat_server/internal/message"
	// "simple_chat_server/internal/model/chatmanager"
)

func main() {

	startApp()
}

func startApp() {
	port := "9000"

	// server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	chatManager := chatmanager.New()

	chatManager.AddGroup(group.New("GENERAL"))

	go chatManager.Run()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go chatmanager.HandleUserConnection(chatManager, conn)
	}
}
