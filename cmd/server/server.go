package main

import (
	"os"
	"simple_chat_server/config"
	"simple_chat_server/internal/chatmanager"
	"simple_chat_server/internal/group"
	"simple_chat_server/internal/model"
	"simple_chat_server/internal/model/logger"

	"log"
	"net"
)

func main() {

	startApp()
}

func startApp() {

	environment := os.Getenv("TKPENV")
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
