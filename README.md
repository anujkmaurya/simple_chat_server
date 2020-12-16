# Simple chat Server

A simple chat application having the following functionalities:
1. The user can join the chat server using Telnet : `telnet ip port`
2. The user can choose a non-existing username for chatting.
3. The user can create a new group and subscribe a group. User can also join an existing group. By default `common` group is created and a new user is automatically assigned this group.
4. The user can leave a group (other than `common` group).
5. The user can choose to ignore messages from a particular user in a group
6. The user can unignore an user, i.e. allow the user again to receive messages fromm that user
7. The user can choose to send a personal message to any person in a group. This message will not be broadcasted and can be used for one to one discussion.
8. Logging all the conversation along with in the defined log file

## Commands Supported 

1. `--joingroup <group name>` : create and join a group 
2. `--leavegroup <group name>` : leave a non-common group
3. `--ignoreuser <user name>` : ignore an user, stop receiving message from that user all groups
4. `--unignoreuser <user name>` : unignore an user, resume  receiving message from that user
5. `--personal <username>` : send personal message to any user

## Code Structure

### Packages
1. `config` :defines the configurations structure
2. `internal/chatmanager` : chatmanager structure and methods. It is responsible for reading user input and performing various actions along with broadcasting message
3. ` internal/command` : command defines the list of alllowed messages/commands, methods to parse messages
4. `internal/group` : contains group structure and it's methods.
5. `internal/message` : contains message structure and it's methods. This message is created to pass user or system message with various metadata info
6. `internal/model` : defines all the constants used by all the packages
7. `internal/model/logger` : reads logger path, creates log file, inits logger 
8. `internal/user` : contains user structure and it's methods.
9. `cmd/server` : contains server startup file. It reads from the config and inits chatmanager package
10. `etc/simple-chat-server` : contains configuration script for each environment

### File Tree

```

.
├── Makefile
├── README.md
├── build
│   └── server
├── chat-server.log
├── cmd
│   └── server
│       └── server.go
├── config
│   ├── config.go
│   ├── init.go
│   └── types.go
├── etc
│   └── simple-chat-server
│       └── development
│           └── simple-char-server.ini
├── go.mod
├── go.sum
└── internal
    ├── chatmanager
    │   ├── init.go
    │   ├── interface.go
    │   ├── manager.go
    │   ├── manager.usecase.go
    │   └── types.go
    ├── command
    │   ├── command.go
    │   └── types.go
    ├── group
    │   ├── group.go
    │   ├── init.go
    │   ├── interface.go
    │   └── types.go
    ├── message
    │   ├── init.go
    │   ├── interface.go
    │   ├── message.go
    │   └── types.go
    ├── model
    │   ├── config.constant.go
    │   ├── constants.go
    │   └── logger
    │       ├── logger.go
    │       └── logger_test.go
    └── user
        ├── init.go
        ├── init_test.go
        ├── interface.go
        ├── types.go
        ├── user.go
        └── user_test.go

```

## Running and building

Compile and Run App using Binary
-----

### Pre-Requisite

1. Golang1.15.4 - Install 'go1.15.4' version from https://golang.org/dl/.
2. checkout the project folder
2. run `make`

## Testing

1. run command `make test` for running testcases and showing detailed test result
2. `make test-short` for  running testcases and showing test result in short
