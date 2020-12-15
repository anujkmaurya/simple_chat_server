package message

import "time"

//IMessage : defines all the exported methods of Message Object
type IMessage interface {
	GetSenderName() string
	GetText() string
	GetReceiverName() string
	SetReceiverName(userName string) IMessage
	GetGroupName() string
	GetCreatedAt() time.Time
	String() string
}
