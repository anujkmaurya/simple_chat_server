package message

import (
	"fmt"
	"time"
)

//GetSenderName : get name of the user who created the message
func (message *Message) GetSenderName() string {
	return message.senderName
}

//GetReceiverName : get name of receipient, if it's empty : the message is not defined for a group
func (message *Message) GetReceiverName() string {
	return message.receiver
}

//GetText : get message text
func (message *Message) GetText() string {
	return message.text
}

//GetGroupName : Get name of the group on which the messsage needs to be send to
func (message *Message) GetGroupName() string {
	return message.groupName
}

//GetCreatedAt : get created at time
func (message *Message) GetCreatedAt() time.Time {
	return message.createdAt
}

//String : gives the message details in string format
func (message *Message) String() string {
	if message.receiver == "" {
		return fmt.Sprintf("{Time:%s, Channel:%s, Sender:%s, Receiver:All, Message:%s}\n", message.createdAt.Format("2006-01-02 15:04:05"), message.groupName, message.senderName, message.text)
	}
	return fmt.Sprintf("{Time:%s, Channel:%s, Sender:%s, Receiver:%s, Message:%s}\n", message.createdAt.Format("2006-01-02 15:04:05"), message.groupName, message.senderName, message.receiver, message.text)
}

//SetReceiverName : assign receiver to the message
func (message *Message) SetReceiverName(userName string) IMessage {
	message.receiver = userName
	return message
}
