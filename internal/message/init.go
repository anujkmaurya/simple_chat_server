package message

import (
	"time"
)

func CreateMessage(sender, channelName, message, receiver string) IMessage {

	return &Message{
		text:        message,
		senderName:  sender,
		channelName: channelName,
		createdAt:   time.Now(),
		receiver:    receiver,
	}

}
