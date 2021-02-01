package girlfriend

var commands = map[string]Command{
	"hug": &DiscordChannelCommand{
		Func: giveHug,
	},
}
