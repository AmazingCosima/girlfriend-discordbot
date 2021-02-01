package girlfriend

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	commandTag = "<@!804384355740680232>"
)

func (app *App) onMessageCreated(session *discordgo.Session, event *discordgo.MessageCreate) {
	content := event.Content
	if !strings.HasPrefix(content, commandTag) {
		return
	}
	app.discordCommand = *event.Message
	content = strings.TrimPrefix(content, commandTag + " ")
	var arguments []string
	arguments = strings.Split(content, " ")
	command := arguments[0]
	if len(arguments) > 1 {
		arguments = arguments[1:]
	} else {
		arguments = nil
	}
	app.RunCommand(command, arguments)
}
