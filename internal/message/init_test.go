package message

import (
	"testing"
)

func TestCreateMessage(t *testing.T) {
	type args struct {
		sender    string
		groupName string
		message   string
		receiver  string
	}
	tests := []struct {
		name           string
		args           args
		wantSenderName string
	}{
		{
			name: "Testcase 1",
			args: args{
				sender:    "sender name",
				groupName: "group name",
				message:   "test message",
				receiver:  "receiver name ",
			},
			wantSenderName: "sender name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMessage(tt.args.sender, tt.args.groupName, tt.args.message, tt.args.receiver); got == nil || got.GetSenderName() != tt.wantSenderName {
				t.Errorf("got.GetSenderName() = %v, wantSenderName %v", got.GetSenderName(), tt.wantSenderName)
			}
		})
	}
}
