package message

import (
	"time"
)

//Message :  format for each chat message
type (
	Message struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}
)
