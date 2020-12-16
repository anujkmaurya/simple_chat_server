package group

import "sync"

//New : create a new Group object
func New(groupName string) IGroup {
	grp := &Group{
		groupName: groupName,
		users:     make(map[string]struct{}, 0),
		mutex:     &sync.RWMutex{},
	}
	return grp
}
