package group

//Group structure: stores the users subscribed to the group
type Group struct {
	groupName string
	users     map[string]struct{}
}
