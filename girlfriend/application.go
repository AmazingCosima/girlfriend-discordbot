package girlfriend

import (
	"github.com/bwmarrin/discordgo"
)

type App struct {
	session *discordgo.Session
}

func (app *App) CreateSession(token string) error{
	var session *discordgo.Session
	var err error
	session, err = discordgo.New("Bot " + token)
	app.session = session
	if err != nil {
		return err
	}
	session.StateEnabled = true
	session.State.MaxMessageCount = 100
	app.session.AddHandler(app.onMessageCreated)
	err = session.Open()
	return err
}