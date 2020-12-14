package message

import "time"

type IMessage interface {
	GetSenderName() string
	GetText() string
	GetReceiverName() string
	GetChannelName() string
	GetCreatedAt() time.Time
	String() string
}
