package group

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		groupName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Testcase 1",
			args: args{
				groupName: "group name",
			},
			want: "group name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.groupName); got == nil || got.GetGroupName() != tt.want {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
