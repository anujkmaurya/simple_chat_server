package chatmanager

//this file contains all teh basic getter setters or Chatmanager struct
import (
	"simple_chat_server/internal/user"
)

//getUser :find if a user is present with given userName
func (um *usersMap) getUser(userName string) (user.IUser, bool) {
	um.mutex.RLock()
	defer um.mutex.RUnlock()

	user, ok := um.users[userName]
	return user, ok
}

//setUser :add user to the user map
func (um *usersMap) setUser(userName string, user user.IUser) {
	um.mutex.Lock()
	defer um.mutex.Unlock()

	um.users[userName] = user
	return
}

//deleteUser :delete user from the user map
func (um *usersMap) deleteUser(userName string) {
	um.mutex.Lock()
	defer um.mutex.Unlock()

	if _, ok := um.users[userName]; ok {
		//delete from usersmap
		delete(um.users, userName)
	}
	return
}
