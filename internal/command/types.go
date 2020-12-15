package command

//Command : specifies allowed user commands
type Command int

//tyoes of commands
const (
	NormalMessage Command = iota + 1
	HelpCommand
	PersonalCommand
	JoinGroupCommand
	LeaveGroupCommand
)

//textToCommandMap converts user command text to defines commands
var textToCommandMap = map[string]Command{
	"--help":       HelpCommand,
	"--personal":   PersonalCommand,
	"--joingroup":  JoinGroupCommand,
	"--leavegroup": LeaveGroupCommand,
}
