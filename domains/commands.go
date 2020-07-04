package domains

type (
	Command func(Context)

	CommandStruct struct {
		command Command
		help    string
	}

	//TODO: New structs for youtube and files
	MemeCommandStruct struct {
		CommandStruct

		youtube string
		//localFile string
		//remoteFile string
	}

	CmdMap     map[string]CommandStruct
	MemeCmdMap map[string]MemeCommandStruct

	CommandHandler struct {
		cmds     CmdMap
		memeCmds MemeCmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap), make(MemeCmdMap)}
}

func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

func (handler CommandHandler) getMemeCmds() MemeCmdMap {
	return handler.memeCmds
}

// Find command inside slice
func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.cmds[name]
	// For legacy reasons, lets just deliver the commnd
	// A new function can be made GetAll() ??
	return &cmd.command, found
}

func (handler CommandHandler) Register(name string, command Command, helpmsg string) {
	// Massage the arguments into a "Full command"
	cmdstruct := CommandStruct{command: command, help: helpmsg}
	handler.cmds[name] = cmdstruct
	if len(name) > 1 {
		handler.cmds[name[:1]] = cmdstruct
	}
}

func (handler CommandHandler) RegisterMemeCmd(name string, cmd Command, helpmsg string, youtubelink string) {
	// Massage the arguments into a "Full command"
	cmdStruct := CommandStruct{command: cmd, help: helpmsg}
	memecmd := MemeCommandStruct{CommandStruct: cmdStruct, youtube: youtubelink}
	handler.memeCmds[name] = memecmd
	if len(name) > 1 {
		handler.memeCmds[name[:1]] = memecmd
	}
}

func (command CommandStruct) GetHelp() string {
	return command.help
}
