package girlfriend

type Command interface {
	RunFunc(app *App, args []string)
}

type BasicCommand struct {
	Func func(app *App, args []string)
}

func (basicCommand *BasicCommand) RunFunc(app *App, args []string) {
	if basicCommand.Func != nil {
		basicCommand.Func(app, args)
	}
}

var commands = map[string]Command{
	"test": &BasicCommand{
		Func:TestCommandFunc,
	},
}

func TestCommandFunc(app *App, args []string) {
	app.SendMessage(app.lastCommandReceived.ChannelID, args[0])
}