package message

import "time"

type IMessage interface {
	GetSenderName() string
	GetText() string
	GetReceiverName() string
	SetReceiverName(userName string) IMessage
	GetChannelName() string
	GetCreatedAt() time.Time
	String() string
}
