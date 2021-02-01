package girlfriend

import (
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	RunFunc(app *App, args []string)
}

type DiscordChannelCommand struct {
	Func func(app *App, message discordgo.Message, args []string)
}

func (channelCommand *DiscordChannelCommand) RunFunc(app *App, args []string) {
	var message discordgo.Message
	message = *app.discordCommand
	app.discordCommand = nil
	if channelCommand.Func != nil {
		channelCommand.Func(app, message, args)
	}
}
