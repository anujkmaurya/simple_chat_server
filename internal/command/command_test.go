package command

import (
	"testing"
)

func TestParseCommand(t *testing.T) {
	type args struct {
		commandList []string
	}

	tests := []struct {
		name string
		args args
		want Command
	}{
		{
			name: "Testcase 1",
			args: args{
				commandList: []string{"some", "text"},
			},
			want: NormalMessage,
		},
		{
			name: "Testcase 2",
			args: args{
				commandList: []string{"--help"},
			},
			want: HelpCommand,
		},
		{
			name: "Testcase 3",
			args: args{
				commandList: []string{"--personal"},
			},
			want: PersonalCommand,
		},
		{
			name: "Testcase 4",
			args: args{
				commandList: []string{"--joingroup"},
			},
			want: JoinGroupCommand,
		},
		{
			name: "Testcase 5",
			args: args{
				commandList: []string{"--leavegroup"},
			},
			want: LeaveGroupCommand,
		},
		{
			name: "Testcase 6",
			args: args{
				commandList: []string{"--ignoreuser"},
			},
			want: IgnoreUserCommand,
		},
		{
			name: "Testcase 7",
			args: args{
				commandList: []string{"--unignoreuser"},
			},
			want: UnIgnoreUserCommand,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommand(tt.args.commandList); got != tt.want {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAndSanitiseCommand(t *testing.T) {
	type args struct {
		commandList []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Testcase 1",
			args: args{
				commandList: []string{"some", "text"},
			},
			want: true,
		},
		{
			name: "Testcase 2",
			args: args{
				commandList: []string{"--help"},
			},
			want: true,
		},
		{
			name: "Testcase 3",
			args: args{
				commandList: []string{"--personal", "user", "message"},
			},
			want: true,
		},
		{
			name: "Testcase 4",
			args: args{
				commandList: []string{"--joingroup", "groupname"},
			},
			want: true,
		},
		{
			name: "Testcase 5",
			args: args{
				commandList: []string{"--leavegroup", "groupname"},
			},
			want: true,
		},
		{
			name: "Testcase 6",
			args: args{
				commandList: []string{"--ignoreuser", "username"},
			},
			want: true,
		},
		{
			name: "Testcase 7",
			args: args{
				commandList: []string{"--unignoreuser", "username"},
			},
			want: true,
		},
		{
			name: "Testcase 8",
			args: args{
				commandList: []string{"--personal", "user"},
			},
			want: false,
		},
		{
			name: "Testcase 9",
			args: args{
				commandList: []string{"--joingroup"},
			},
			want: false,
		},
		{
			name: "Testcase 10",
			args: args{
				commandList: []string{"--leavegroup"},
			},
			want: false,
		},
		{
			name: "Testcase 11",
			args: args{
				commandList: []string{"--ignoreuser"},
			},
			want: false,
		},
		{
			name: "Testcase 12",
			args: args{
				commandList: []string{"--unignoreuser"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseAndSanitiseCommand(tt.args.commandList); got != tt.want {
				t.Errorf("ParseAndSanitiseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
