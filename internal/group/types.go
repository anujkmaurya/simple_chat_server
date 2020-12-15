package group

//group structure
type Group struct {
	groupName string
	users     map[string]struct{}
}
