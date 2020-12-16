package group

import "sync"

//Group structure: stores the users subscribed to the group
type Group struct {
	mutex     *sync.RWMutex
	groupName string
	users     map[string]struct{}
}
