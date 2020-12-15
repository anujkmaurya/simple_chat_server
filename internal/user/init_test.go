package user

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Testcase 1",
			args: args{
				name: "newUser",
			},
			want: "newUser",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name); got.GetUserName() != tt.want {
				t.Errorf("New() = %v, want %v", got.GetUserName(), tt.want)
			}
		})
	}
}
