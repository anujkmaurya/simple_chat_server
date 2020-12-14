package message

import (
	"time"
)

func CreateMessage(sender, channelName, message string) IMessage {

	return &Message{
		text:        message,
		senderName:  sender,
		channelName: channelName,
		createdAt:   time.Now(),
	}

}
