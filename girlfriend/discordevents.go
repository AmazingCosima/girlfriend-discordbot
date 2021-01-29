package girlfriend

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (app *App) onMessageCreated(session *discordgo.Session, event *discordgo.MessageCreate) {
	channelID := event.ChannelID
	content := event.Content
	if !strings.HasPrefix(content, "gf ") {
		return
	}
	var arguments []string
	content = strings.TrimPrefix(content, "gf ")
	arguments = strings.Split(content, " ")
	app.SendMessage(channelID, arguments[0])
}
