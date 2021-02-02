package girlfriend

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

var hugs = []string{
	"https://tenor.com/view/anime-hug-gif-15249774",
	"https://tenor.com/view/anime-hug-love-gif-15900664",
	"https://tenor.com/view/anime-hug-hug-hugs-anime-girl-anime-girl-hug-gif-16787485",
	"https://tenor.com/view/hug-anime-sweet-couple-gif-12668480",
	"https://tenor.com/view/hug-anime-gif-4898650",
	"https://tenor.com/view/hug-love-sweet-anime-couple-gif-16300141",
	"https://tenor.com/view/sao-hug-anime-sword-art-online-asuna-gif-16993945",
}

func giveHug(app *App, message discordgo.Message, args []string) {
	var hug string
	hug = hugs[rand.Intn(len(hugs))]
	app.SendMessage(message.ChannelID, hug)
	app.SendMessage(message.ChannelID, args[0])
}
