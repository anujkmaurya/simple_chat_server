package group

func New(groupName string) IGroup {
	grp := &Group{
		groupName: groupName,
		users:     make(map[string]string, 0),
	}
	return grp
}
