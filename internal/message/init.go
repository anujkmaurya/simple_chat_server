package message

import (
	"time"
)

//CreateMessage : Create New Message
func CreateMessage(sender, groupName, message, receiver string) IMessage {

	return &Message{
		text:       message,
		senderName: sender,
		groupName:  groupName,
		createdAt:  time.Now(),
		receiver:   receiver,
	}

}
