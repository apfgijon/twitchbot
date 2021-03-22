package twitchbot

import (
	"github.com/apfgijon/cartones/internal/pkg/client"
	"github.com/apfgijon/cartones/internal/pkg/generalbot"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
	"github.com/gempir/go-twitch-irc/v2"
)

type Twitchbot struct {
	botName string
	channel string
	oauth   string
}

func (t *Twitchbot) Build(botName string, channel string, oauth string) {
	t.botName = botName
	t.channel = channel
	t.oauth = oauth
}

func (t *Twitchbot) Bootstrap() {

	client := client.Communication{
		Client:  twitch.NewClient(t.botName, t.oauth),
		Channel: t.channel,
		BotName: t.botName,
	}

	poke := &pokemon.PokeapiImpl{}

	generalbot := &generalbot.Generalbot{}

	covid := &covid.CovidApiImpl{}

	generalbot.Init(poke, client, covid)
	generalbot.Start()

}
