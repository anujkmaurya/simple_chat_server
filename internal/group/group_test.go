package group

import (
	"sync"
	"testing"
)

func TestGroup_SetGroupName(t *testing.T) {
	type fields struct {
		mutex     *sync.RWMutex
		groupName string
		users     map[string]struct{}
	}
	type args struct {
		groupName string
	}
	mutex := &sync.RWMutex{}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex:     mutex,
				groupName: "test group",
				users:     make(map[string]struct{}, 0),
			},
			args: args{
				groupName: "new group name",
			},
			want: "new group name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				mutex:     tt.fields.mutex,
				groupName: tt.fields.groupName,
				users:     tt.fields.users,
			}
			g.SetGroupName(tt.args.groupName)
			if got := g.GetGroupName(); got != tt.want {
				t.Errorf("SetGroupName(), got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_GetGroupName(t *testing.T) {
	type fields struct {
		mutex     *sync.RWMutex
		groupName string
		users     map[string]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex:     &sync.RWMutex{},
				groupName: "test group",
				users:     make(map[string]struct{}, 0),
			},
			want: "test group",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				mutex:     tt.fields.mutex,
				groupName: tt.fields.groupName,
				users:     tt.fields.users,
			}
			if got := g.GetGroupName(); got != tt.want {
				t.Errorf("Group.GetGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_AddUserToGroup(t *testing.T) {
	type fields struct {
		mutex     *sync.RWMutex
		groupName string
		users     map[string]struct{}
	}
	type args struct {
		userName string
	}

	users := make(map[string]struct{}, 0)
	users["old user"] = struct{}{}

	tests := []struct {
		name         string
		fields       fields
		args         args
		want         bool
		wantUserName string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex:     &sync.RWMutex{},
				groupName: "test group",
				users:     make(map[string]struct{}, 0),
			},
			args: args{
				userName: "new user",
			},
			want:         true,
			wantUserName: "new user",
		},
		{
			name: "Testcase 2",
			fields: fields{
				mutex:     &sync.RWMutex{},
				groupName: "test group",
				users:     users,
			},
			args: args{
				userName: "old user",
			},
			want:         false,
			wantUserName: "old user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				mutex:     tt.fields.mutex,
				groupName: tt.fields.groupName,
				users:     tt.fields.users,
			}
			if got := g.AddUserToGroup(tt.args.userName); got != tt.want {
				t.Errorf("Group.AddUserToGroup() = %v, want %v", got, tt.want)
			}
			found := false
			for k := range g.GetSubscribedUsers() {
				if k == tt.wantUserName {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Failed: user not added to user map")
			}
		})
	}
}

func TestGroup_RemoveUserFromGroup(t *testing.T) {
	type fields struct {
		mutex     *sync.RWMutex
		groupName string
		users     map[string]struct{}
	}

	users := make(map[string]struct{}, 0)
	users["old user"] = struct{}{}

	type args struct {
		userName string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantUserName string
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex:     &sync.RWMutex{},
				groupName: "test group",
				users:     users,
			},
			args: args{
				userName: "old user",
			},
			wantUserName: "old user",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				mutex:     tt.fields.mutex,
				groupName: tt.fields.groupName,
				users:     tt.fields.users,
			}
			g.RemoveUserFromGroup(tt.args.userName)
			found := false
			for k := range g.GetSubscribedUsers() {
				if k == tt.wantUserName {
					found = true
					break
				}
			}
			if found {
				t.Errorf("Failed: user still present in user map")
			}
		})
	}
}

func TestGroup_GetSubscribedUsersCount(t *testing.T) {
	type fields struct {
		mutex     *sync.RWMutex
		groupName string
		users     map[string]struct{}
	}

	users := make(map[string]struct{}, 0)
	users["old user"] = struct{}{}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Testcase 1",
			fields: fields{
				mutex:     &sync.RWMutex{},
				groupName: "test group",
				users:     make(map[string]struct{}, 0),
			},

			want: 0,
		},
		{
			name: "Testcase 2",
			fields: fields{
				mutex:     &sync.RWMutex{},
				groupName: "test group",
				users:     users,
			},

			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				mutex:     tt.fields.mutex,
				groupName: tt.fields.groupName,
				users:     tt.fields.users,
			}
			if got := g.GetSubscribedUsersCount(); got != tt.want {
				t.Errorf("Group.GetSubscribedUsersCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
