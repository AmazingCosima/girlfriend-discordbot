package girlfriend

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (app *App) SendMessage(channelID, content string) {
	message, err := app.session.ChannelMessageSend(channelID, content)
	if err != nil {
		return
	}
	app.logger.Log("Created message " + message.ID + " in " + channelID)
}

func (app *App) DeleteMessage(channelID, messageID string) {
	err := app.session.ChannelMessageDelete(channelID, messageID)
	if err != nil {
		return
	}
	app.logger.Log("Destroyed message " + messageID + " in " + channelID)
}

func (app *App) ParseFromMention(mention string) *discordgo.User{
	if !strings.HasPrefix(mention, "<@!") && !strings.HasSuffix(mention, ">") {
		return nil
	}
	userID := strings.TrimSuffix(strings.TrimPrefix(mention, "<@!"), ">")
	user, err := app.session.User(userID)
	if err != nil {
		return nil
	}
	return user
}

func (app *App) ParseMention(user *discordgo.User) string{
	var mention string
	if user == nil {
		return mention
	}
	userID := user.ID
	mention = "<@!" + userID + ">"
	return mention
}
