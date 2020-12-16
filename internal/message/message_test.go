package message

import (
	"reflect"
	"testing"
	"time"
)

func TestMessage_GetSenderName(t *testing.T) {
	type fields struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  time.Now(),
				receiver:   "test receiver",
			},
			want: "sender",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				text:       tt.fields.text,
				senderName: tt.fields.senderName,
				groupName:  tt.fields.groupName,
				createdAt:  tt.fields.createdAt,
				receiver:   tt.fields.receiver,
			}
			if got := message.GetSenderName(); got != tt.want {
				t.Errorf("Message.GetSenderName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetReceiverName(t *testing.T) {
	type fields struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  time.Now(),
				receiver:   "test receiver",
			},
			want: "test receiver",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				text:       tt.fields.text,
				senderName: tt.fields.senderName,
				groupName:  tt.fields.groupName,
				createdAt:  tt.fields.createdAt,
				receiver:   tt.fields.receiver,
			}
			if got := message.GetReceiverName(); got != tt.want {
				t.Errorf("Message.GetReceiverName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetText(t *testing.T) {
	type fields struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  time.Now(),
				receiver:   "test receiver",
			},
			want: "test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				text:       tt.fields.text,
				senderName: tt.fields.senderName,
				groupName:  tt.fields.groupName,
				createdAt:  tt.fields.createdAt,
				receiver:   tt.fields.receiver,
			}
			if got := message.GetText(); got != tt.want {
				t.Errorf("Message.GetText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetGroupName(t *testing.T) {
	type fields struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  time.Now(),
				receiver:   "test receiver",
			},
			want: "test group",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				text:       tt.fields.text,
				senderName: tt.fields.senderName,
				groupName:  tt.fields.groupName,
				createdAt:  tt.fields.createdAt,
				receiver:   tt.fields.receiver,
			}
			if got := message.GetGroupName(); got != tt.want {
				t.Errorf("Message.GetGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_GetCreatedAt(t *testing.T) {
	type fields struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}

	currTime := time.Now()

	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name: "Testcase 1",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  currTime,
				receiver:   "test receiver",
			},
			want: currTime,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				text:       tt.fields.text,
				senderName: tt.fields.senderName,
				groupName:  tt.fields.groupName,
				createdAt:  tt.fields.createdAt,
				receiver:   tt.fields.receiver,
			}
			if got := message.GetCreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Message.GetCreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessage_String(t *testing.T) {
	type fields struct {
		text       string
		senderName string
		groupName  string
		createdAt  time.Time
		receiver   string
	}

	currTime := time.Now()

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1- receiver present",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  currTime,
				receiver:   "test receiver",
			},
			want: "{Time:" + currTime.Format("2006-01-02 15:04:05") + ", Channel:test group, Sender:sender, Receiver:test receiver, Message:test message}\n",
		},
		{
			name: "Testcase 2 - for all users, receiver empty",
			fields: fields{
				text:       "test message",
				senderName: "sender",
				groupName:  "test group",
				createdAt:  currTime,
				receiver:   "",
			},
			want: "{Time:" + currTime.Format("2006-01-02 15:04:05") + ", Channel:test group, Sender:sender, Receiver:All, Message:test message}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				text:       tt.fields.text,
				senderName: tt.fields.senderName,
				groupName:  tt.fields.groupName,
				createdAt:  tt.fields.createdAt,
				receiver:   tt.fields.receiver,
			}
			if got := message.String(); got != tt.want {
				t.Errorf("Message.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
