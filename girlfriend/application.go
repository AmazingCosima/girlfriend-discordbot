package girlfriend

import (
	"github.com/bwmarrin/discordgo"
)

type App struct {
	session *discordgo.Session
	logger *Logger
}

func (app *App) CreateSession(token string) error{
	var session *discordgo.Session
	var err error
	session, err = discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	session.StateEnabled = true
	session.State.MaxMessageCount = 100
	app.session = session
	app.logger = new(Logger)
	app.session.AddHandler(app.onMessageCreated)
	err = session.Open()
	if err == nil {
		app.logger.Log("Successfully created session")
	}
	return err
}
