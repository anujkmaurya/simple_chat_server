package message

import (
	"fmt"
	"time"
)

func (message *Message) GetSenderName() string {
	return message.senderName
}

func (message *Message) GetReceiverName() string {
	return message.receiver
}

func (message *Message) GetText() string {
	return message.text
}

func (message *Message) GetChannelName() string {
	return message.channelName
}

func (message *Message) GetCreatedAt() time.Time {
	return message.createdAt
}

func (message *Message) String() string {
	return fmt.Sprintf("{Time:%s, Channel:%s, Sender:%s, Message:%s}\n", message.createdAt.Format("2006-01-02 15:04:05"), message.channelName, message.senderName, message.text)
}
