package girlfriend

func (app *App) SendMessage(channelID, content string) {
	message, err := app.session.ChannelMessageSend(channelID, content)
	if err != nil {
		return
	}
	app.logger.Log("Created message " + message.ID)
}
