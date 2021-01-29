package girlfriend

func (app *App) SendMessage(channelID, content string) {
	message, err := app.session.ChannelMessageSend(channelID, content)
	if err != nil {
	}
	_ = message
}