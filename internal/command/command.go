package command

//ParseCommand : parse input and decides command type
func ParseCommand(commandList []string) Command {

	if cmd, ok := textToCommandMap[commandList[0]]; ok {
		return cmd
	}
	return NormalMessage

}

//ParseAndSanitiseCommand : sanitise the input, commands received from user
func ParseAndSanitiseCommand(commandList []string) bool {
	cmd := ParseCommand(commandList)

	cmdListLen := len(commandList)

	switch cmd {
	case HelpCommand:
		return true

	case PersonalCommand:
		if cmdListLen > 2 {
			return true
		}
	case JoinGroupCommand, LeaveGroupCommand:
		if cmdListLen == 2 {
			return true
		}
	case NormalMessage:
		return true
	}

	return false
}
