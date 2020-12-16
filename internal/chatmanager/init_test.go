package chatmanager

import (
	"simple_chat_server/internal/user"
	"testing"
)

func TestNew(t *testing.T) {
	type fields struct {
		userName string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				userName: "test user",
			},
			want: "test user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New()
			got.AddUser(user.New(tt.fields.userName))

			if _, err := got.GetUser(tt.want); err != nil {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
