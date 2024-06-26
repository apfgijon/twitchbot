package twitchbot

import (
	"github.com/apfgijon/cartones/internal/pkg/A-comunication/bot"
	"github.com/apfgijon/cartones/internal/pkg/A-comunication/client"
	"github.com/apfgijon/cartones/internal/pkg/services"
	"github.com/gempir/go-twitch-irc/v2"
)

type Twitchbot struct {
	botName    string
	channel    string
	oauth      string
	pokeGame   string
	generalbot bot.Bot
}

func NewTwitchBot(botName string, channel string, oauth string, pokeGame string) *Twitchbot {
	return &Twitchbot{
		botName:  botName,
		channel:  channel,
		oauth:    oauth,
		pokeGame: pokeGame,
	}
}

func (t *Twitchbot) Bootstrap() {

	client := client.Communication{
		Client:  twitch.NewClient(t.botName, t.oauth),
		Channel: t.channel,
		BotName: t.botName,
	}

	t.generalbot, _ = services.InitializeBot(client, t.pokeGame)

	t.generalbot.Start()

}
func (t *Twitchbot) Stop() {
	t.generalbot.Stop()
}
