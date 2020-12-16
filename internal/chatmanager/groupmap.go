package chatmanager

//this file contains all teh basic getter setters or Chatmanager struct
import (
	"simple_chat_server/internal/group"
)

//getGroup :find if a group is present with given groupName
func (gm *groupMap) getGroup(groupName string) (group.IGroup, bool) {
	gm.mutex.RLock()
	defer gm.mutex.RUnlock()

	group, ok := gm.groups[groupName]
	return group, ok
}

//setUser: add a group to the group map
func (gm *groupMap) setGroup(groupName string, group group.IGroup) {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	gm.groups[groupName] = group
	return
}

//deleteGroup :delete group from the group map
func (gm *groupMap) deleteGroup(groupName string) {
	gm.mutex.Lock()
	defer gm.mutex.Unlock()

	if _, ok := gm.groups[groupName]; ok {
		//delete from groupmap
		delete(gm.groups, groupName)
	}
	return
}
