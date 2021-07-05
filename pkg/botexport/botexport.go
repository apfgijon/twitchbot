package botexport

import "github.com/apfgijon/cartones/internal/twitchbot"

func InitBot(botname string, channel string, oauth string, game string) {
	t := twitchbot.NewTwitchBot(botname, channel, oauth, game)
	t.Bootstrap()
}
