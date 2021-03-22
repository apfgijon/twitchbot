package generalbot

import (
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/client"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
	"github.com/apfgijon/cartones/pkg/randomsay"
	"github.com/gempir/go-twitch-irc/v2"
)

type Bot interface {
	Init(p pokemon.PokeInfo, c client.Communication)
	Start()
}

type Generalbot struct {
	com   client.Communication
	poke  pokemon.PokeInfo
	covid covid.CovidInfo
}

func (gn *Generalbot) Init(p pokemon.PokeInfo, c client.Communication, cov covid.CovidInfo) {
	gn.poke = p
	gn.com = c
	gn.covid = cov
}

func (gn *Generalbot) Start() {
	gn.poke.Build()

	gn.com.Client.OnPrivateMessage(gn.onMessage)

	gn.com.Client.Join(gn.com.Channel)
	// go gn.sayRandomPhrase()
	// go gn.sayRandomRefran()
	err := gn.com.Client.Connect()
	if err != nil {
		panic(err)
	}
}

func (gn *Generalbot) onMessage(message twitch.PrivateMessage) {
	if !gn.checkCommands(message) {
		if strings.Contains(strings.ToUpper(message.Message), strings.ToUpper(gn.com.BotName)) {
			message := "Que me dices " + message.User.DisplayName + "? nun ves que soy un bot? Amás nun pescancio castel.lán."
			gn.com.Client.Say(gn.com.Channel, message)
		}
		if string(message.Message[0]) != "!" && string(message.Message[0]) != "@" && message.User.DisplayName != "Nightbot" && !strings.Contains(strings.ToLower(message.Message), "zonnyo") {
			randomsay.SetPhrase(message.Message)
		}
	}
}
