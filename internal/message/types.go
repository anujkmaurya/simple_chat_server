package message

import (
	"time"
)

//Message format
type (
	Message struct {
		text        string
		senderName  string
		channelName string
		createdAt   time.Time
		receiver    string
	}
)
