package girlfriend

import (
	"github.com/bwmarrin/discordgo"
)

func Login(token string) error{
	var session *discordgo.Session
	var err error
	session, err = discordgo.New(token)
	if err != nil {
		return err
	}
	err = session.Open()
	return err
}