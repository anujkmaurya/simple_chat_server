package group

func New(groupName string) IGroup {
	grp := &Group{
		groupName: groupName,
		users:     make(map[string]struct{}, 0),
	}
	return grp
}
