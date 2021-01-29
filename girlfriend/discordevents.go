package girlfriend

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	commandTag = "gf"
)

func (app *App) onMessageCreated(session *discordgo.Session, event *discordgo.MessageCreate) {
	channelID := event.ChannelID
	content := event.Content
	if !strings.HasPrefix(content, commandTag) {
		return
	}
	var arguments []string
	content = strings.TrimPrefix(content, commandTag + " ")
	arguments = strings.Split(content, " ")
	app.SendMessage(channelID, arguments[0])
}
