package chatmanager

import (
	"simple_chat_server/internal/user"
	"sync"
	"testing"
)

func Test_usersMap_getUser(t *testing.T) {
	type fields struct {
		mutex *sync.RWMutex
		users map[string]user.IUser
	}

	tUser := make(map[string]user.IUser, 0)
	tUser["test user"] = user.New("test user")

	type args struct {
		userName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex: &sync.RWMutex{},
				users: tUser,
			},
			args: args{
				userName: "test user",
			},
			want:  "test user",
			want1: true,
		},
		{
			name: "Testcase 2",
			fields: fields{
				mutex: &sync.RWMutex{},
				users: make(map[string]user.IUser, 0),
			},
			args: args{
				userName: "test user",
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			um := &usersMap{
				mutex: tt.fields.mutex,
				users: tt.fields.users,
			}
			got, got1 := um.getUser(tt.args.userName)
			if got1 != tt.want1 {
				t.Errorf("usersMap.getUser() got1 = %v, want %v", got1, tt.want1)
				return
			}
			if got != nil && got.GetUserName() != tt.want {
				t.Errorf("usersMap.getUser() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_usersMap_setUser(t *testing.T) {
	type fields struct {
		mutex *sync.RWMutex
		users map[string]user.IUser
	}

	tUser := make(map[string]user.IUser, 0)

	type args struct {
		userName string
		user     user.IUser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex: &sync.RWMutex{},
				users: tUser,
			},
			args: args{
				userName: "test user",
			},
			want:  "test user",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			um := &usersMap{
				mutex: tt.fields.mutex,
				users: tt.fields.users,
			}
			um.setUser(tt.args.userName, tt.args.user)

			got, got1 := um.getUser(tt.args.userName)
			if got1 != tt.want1 {
				t.Errorf("usersMap.setUser() got1 = %v, want %v", got1, tt.want1)
				return
			}
			if got != nil && got.GetUserName() != tt.want {
				t.Errorf("usersMap.setUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usersMap_deleteUser(t *testing.T) {
	type fields struct {
		mutex *sync.RWMutex
		users map[string]user.IUser
	}

	tUser := make(map[string]user.IUser, 0)
	tUser["test user"] = user.New("test user")

	type args struct {
		userName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex: &sync.RWMutex{},
				users: tUser,
			},
			args: args{
				userName: "test user",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			um := &usersMap{
				mutex: tt.fields.mutex,
				users: tt.fields.users,
			}
			um.deleteUser(tt.args.userName)
			_, got1 := um.getUser(tt.args.userName)
			if got1 != tt.want {
				t.Errorf("usersMap.deleteUser() failed,  got1 = %v, want %v", got1, tt.want)
				return
			}
		})
	}
}
