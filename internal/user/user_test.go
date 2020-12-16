package user

import (
	"reflect"
	"simple_chat_server/internal/message"
	"simple_chat_server/internal/model"
	"sync"
	"testing"
)

func TestUser_GetUserName(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
		mutex        *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
				currentGroup: model.CommonGroup,
				groups:       make(map[string]struct{}, 0),
				mutex:        &sync.RWMutex{},
			},
			want: "testUser",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
				mutex:        tt.fields.mutex,
			}
			if got := u.GetUserName(); got != tt.want {
				t.Errorf("User.GetUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetOutChannel(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
		mutex        *sync.RWMutex
	}
	outChannel := make(chan message.IMessage, model.MaxUserMessageQueueLen)
	tests := []struct {
		name   string
		fields fields
		want   chan message.IMessage
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          outChannel,
				currentGroup: model.CommonGroup,
				groups:       make(map[string]struct{}, 0),
				mutex:        &sync.RWMutex{},
			},
			want: outChannel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
				mutex:        tt.fields.mutex,
			}
			if got := u.GetOutChannel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.GetOutChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SendMessageToUser(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
		mutex        *sync.RWMutex
	}
	type args struct {
		message message.IMessage
	}

	msg := message.CreateMessage("sender", "common", "test message", "abc")

	tests := []struct {
		name   string
		fields fields
		args   args
		want   message.IMessage
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
				currentGroup: model.CommonGroup,
				groups:       make(map[string]struct{}, 0),
				mutex:        &sync.RWMutex{},
			},
			args: args{
				message: msg,
			},
			want: msg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
				mutex:        tt.fields.mutex,
			}
			go u.SendMessageToUser(tt.args.message)

			if got, ok := <-u.out; ok && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.SendMessageToUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetCurrentUserGroup(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
		mutex        *sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
				currentGroup: model.CommonGroup,
				groups:       make(map[string]struct{}, 0),
				mutex:        &sync.RWMutex{},
			},
			want: model.CommonGroup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
				mutex:        tt.fields.mutex,
			}
			if got := u.GetCurrentUserGroup(); got != tt.want {
				t.Errorf("User.GetCurrentUserGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetCurrentUserGroup(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
		mutex        *sync.RWMutex
	}
	type args struct {
		groupName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
				currentGroup: "",
				groups:       make(map[string]struct{}, 0),
				mutex:        &sync.RWMutex{},
			},
			args: args{
				groupName: model.CommonGroup,
			},
			want: model.CommonGroup,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
				mutex:        tt.fields.mutex,
			}
			u.SetCurrentUserGroup(tt.args.groupName)
			if got := u.GetCurrentUserGroup(); got != tt.want {
				t.Errorf("User.SetCurrentUserGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_GetIgnoredUserName(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
				currentGroup: model.CommonGroup,
				groups:       make(map[string]struct{}, 0),
				ignoredUser:  "ignoredUser",
			},
			want: "ignoredUser",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
			}
			if got := u.GetIgnoredUserName(); got != tt.want {
				t.Errorf("User.GetIgnoredUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetIgnoredUserName(t *testing.T) {
	type fields struct {
		name         string
		out          chan message.IMessage
		currentGroup string
		ignoredUser  string
		groups       map[string]struct{}
	}
	type args struct {
		userName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				name:         "testUser",
				out:          make(chan message.IMessage, model.MaxUserMessageQueueLen),
				currentGroup: model.CommonGroup,
				groups:       make(map[string]struct{}, 0),
				ignoredUser:  "",
			},
			args: args{
				userName: "ignoredUser",
			},
			want: "ignoredUser",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				name:         tt.fields.name,
				out:          tt.fields.out,
				currentGroup: tt.fields.currentGroup,
				ignoredUser:  tt.fields.ignoredUser,
				groups:       tt.fields.groups,
			}
			u.SetIgnoredUserName(tt.args.userName)
			if got := u.GetIgnoredUserName(); got != tt.want {
				t.Errorf("User.SetIgnoredUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
