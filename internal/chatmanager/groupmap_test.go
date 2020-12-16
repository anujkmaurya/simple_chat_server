package chatmanager

import (
	"simple_chat_server/internal/group"
	"sync"
	"testing"
)

func Test_groupMap_getGroup(t *testing.T) {
	type fields struct {
		mutex  *sync.RWMutex
		groups map[string]group.IGroup
	}
	type args struct {
		groupName string
	}

	tGroup := make(map[string]group.IGroup, 0)
	tGroup["test group"] = group.New("test group")

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
				mutex:  &sync.RWMutex{},
				groups: make(map[string]group.IGroup, 0),
			},
			args: args{
				groupName: "test group",
			},
			want:  "",
			want1: false,
		},
		{
			name: "Testcase 2",
			fields: fields{
				mutex:  &sync.RWMutex{},
				groups: tGroup,
			},
			args: args{
				groupName: "test group",
			},
			want:  "test group",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &groupMap{
				mutex:  tt.fields.mutex,
				groups: tt.fields.groups,
			}
			got, got1 := gm.getGroup(tt.args.groupName)
			if got1 != tt.want1 {
				t.Errorf("groupMap.getGroup() got1 = %v, want %v", got1, tt.want1)
				return
			}
			if got != nil && got.GetGroupName() != tt.want {
				t.Errorf("groupMap.getGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupMap_setGroup(t *testing.T) {
	type fields struct {
		mutex  *sync.RWMutex
		groups map[string]group.IGroup
	}

	tGroup := make(map[string]group.IGroup, 0)

	type args struct {
		groupName string
		group     group.IGroup
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
				mutex:  &sync.RWMutex{},
				groups: tGroup,
			},
			args: args{
				groupName: "test group",
			},
			want:  "test group",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &groupMap{
				mutex:  tt.fields.mutex,
				groups: tt.fields.groups,
			}
			gm.setGroup(tt.args.groupName, tt.args.group)

			got, got1 := gm.getGroup(tt.args.groupName)
			if got1 != tt.want1 {
				t.Errorf("groupMap.setGroup() got1 = %v, want %v", got1, tt.want1)
				return
			}
			if got != nil && got.GetGroupName() != tt.want {
				t.Errorf("groupMap.setGroup() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupMap_deleteGroup(t *testing.T) {
	type fields struct {
		mutex  *sync.RWMutex
		groups map[string]group.IGroup
	}

	tGroup := make(map[string]group.IGroup, 0)
	tGroup["test group"] = group.New("test group")

	type args struct {
		groupName string
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
				mutex:  &sync.RWMutex{},
				groups: tGroup,
			},
			args: args{
				groupName: "test group",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gm := &groupMap{
				mutex:  tt.fields.mutex,
				groups: tt.fields.groups,
			}
			gm.deleteGroup(tt.args.groupName)
			_, got1 := gm.getGroup(tt.args.groupName)
			if got1 != tt.want {
				t.Errorf("groupMap.deleteGroup() failed,  got1 = %v, want %v", got1, tt.want)
				return
			}
		})
	}
}
