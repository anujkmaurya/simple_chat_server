# Simple chat Server

A simpe chat application purely over TCP

## Code Structure

### Packages
1. `config` :defines the configurations structure
2. `internal/chatmanager` : chatmanager structure and methods. It is responsible for reading user input and performing various actions along with broadcasting message
3.` internal/command` : command defines the list of alllowed messages/commands, methods to parse messages
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

### Important Points 



## Running and building

Compile and Run App using Binary
-----

### Pre-Requisite

1. Golang1.15.4 - Install 'go1.15.4' version from https://golang.org/dl/.

### Steps:

1. Create folder  `github.com/personal-work/` in `~/go/src/` directory in your laptop.
2. clone or download the files
3. extract and move the folder `message_delivery_sys` to `~/go/src/github.com/personal-work/` directory
4. go to `~/go/src/github.com/personal-work/message_delivery_sys`
5. resolve dependencies using `dep ensure -v`
6. Compile and build using command `make build`
7. go to `build` directory and Run the binary file `./server` on one terminal
8. go to `build` directory and Run the binary file `./client` on another terminal
9. repeat step 8, to spawn as many client limited to 255 


## Testing

1. follow all the ![Pre-Requisite] and steps 1-5 of ![Steps]
2. run command `make test`
