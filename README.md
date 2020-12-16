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

sample logs 
```
2020/12/16 14:58:43 {Time:2020-12-16 14:58:43, Channel:COMMON, Sender:SYSTEM, Receiver:All, Message:anuj has joined the channel. Say hello.}
2020/12/16 14:58:56 {Time:2020-12-16 14:58:56, Channel:COMMON, Sender:SYSTEM, Receiver:All, Message:vivek has joined the channel. Say hello.}
2020/12/16 14:58:59 {Time:2020-12-16 14:58:59, Channel:COMMON, Sender:vivek, Receiver:All, Message:hi anuj}
2020/12/16 14:59:05 {Time:2020-12-16 14:59:05, Channel:COMMON, Sender:anuj, Receiver:All, Message:hi vivek}

```

## Commands Supported 

1. `--joingroup <group name>` : create and join a group 
2. `--leavegroup <group name>` : leave a non-common group
3. `--ignoreuser <user name>` : ignore an user, stop receiving message from that user all groups
4. `--unignoreuser <user name>` : unignore an user, resume  receiving message from that user
5. `--personal <username>` : send personal message to any user

## Code Structure

### Packages
1. `config` :defines the configurations structure for chat server
2. `internal/chatmanager` : chatmanager structure and methods. It is responsible for reading user input and performing various actions along with broadcasting message
3. ` internal/command` : command defines the list of alllowed messages/commands, methods to parse messages
4. `internal/group` : contains group structure and it's methods.
5. `internal/message` : contains message structure and it's methods. This message is created to pass user or system message with various metadata info
6. `internal/model` : defines all the constants used by all the packages
7. `internal/model/logger` : reads logger path, creates log file, inits logger 
8. `internal/user` : contains user structure and it's methods.
9. `cmd/server` : contains server startup file. It reads from the config and inits chatmanager package
10. `etc/simple-chat-server` : contains configuration script for each environment
11. `internal/mock` : contains the mock config, mock log files

### File Tree

```

.
├── Dockerfile
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
│   ├── config_test.go
│   ├── init.go
│   ├── init_test.go
│   └── types.go
├── etc
│   └── simple-chat-server
│       └── development
│           └── simple-chat-server.ini
├── go.mod
├── go.sum
└── internal
    ├── chatmanager
    │   ├── groupmap.go
    │   ├── groupmap_test.go
    │   ├── init.go
    │   ├── init_test.go
    │   ├── interface.go
    │   ├── manager.go
    │   ├── manager.usecase.go
    │   ├── types.go
    │   ├── usermap.go
    │   └── usermap_test.go
    ├── command
    │   ├── command.go
    │   ├── command_test.go
    │   └── types.go
    ├── group
    │   ├── group.go
    │   ├── group_test.go
    │   ├── init.go
    │   ├── init_test.go
    │   ├── interface.go
    │   └── types.go
    ├── message
    │   ├── init.go
    │   ├── init_test.go
    │   ├── interface.go
    │   ├── message.go
    │   ├── message_test.go
    │   └── types.go
    ├── mock
    │   ├── config
    │   │   └── config-mock.ini
    │   └── files
    │       └── testfile
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

## Compile and build server binary
 run `make`

## Testing

1. run command `make test` for running testcases and showing detailed test result
2. `make test-short` for  running testcases and showing test result in short


### Running with Docker

extract simple_chat_server.zip and navigate to simple_chat_server folder 

#### To build the image from the Dockerfile, run:

`docker build -t chat-server .`

#### To start an interactive shell, run:

`docker run -p 9000:9000 -it --rm --name chat-server chat-server`

## Extra Features

1. Group Wise chat : users can join a group. If the group is not created, it will created first and then user will be switched to this new group. Users can choose to leave group also, on leaving user will be to switched to "common" group.
2. Clean up of Group: If all the users in a group leave a group and that group is not "common", it will be deleted.
3. Client Disconnect handled: If a user is disconnected from the session, he will be unsubscribed from all the groups he subscribed to and then the user will be deleted. So a new user can join with the same name and can start fresh.
4. System errors along with all chats are logged in the log file.
5. The chat message are of the JSON format, so they can be read and parsed easily for analytics purpose.
6. One can choose to block/unblock a particular user present in group.
7. Unit test cases added for majority of the methods and packages.

## Limitation and Possible Enhancements

1. No Check for max users. We can define max users that should be able to connect to this server. It can be configured from simple-char-server.ini which can be checked when a user joins.
2. We can fine tune the TCP connection with clients by setting various Read/Write timeout from config file.
3. At present all the chats along with system errors are logged in the logfile. We can define a separate file for logging system errors.
4. We can send the chat message which are already in the JSON format to a analyser and dunmp it in a time series database. We can then perform various analytics on it.
5. We can persist the message as mentioned in point #4, it can be shown as history.
6. We can add more testcases, write benchmark test and integration testcase.
7. We can spin a http server from main in the go routine and expose various endpoints to make various announcements as a system : POST /groups/{group_name}
8. Add authentication before allowing user to start chat session.


