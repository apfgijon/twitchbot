package twitchbot

import (
	"github.com/apfgijon/cartones/internal/pkg/client"
	"github.com/apfgijon/cartones/internal/pkg/generalbot"
	"github.com/gempir/go-twitch-irc/v2"
)

func Start(botName string, channel string, oauth string) {

	client := client.Communication{
		Client:  twitch.NewClient(botName, oauth),
		Channel: channel,
		BotName: botName,
	}

	generalbot := generalbot.Generalbot{
		Com: client,
	}

	generalbot.Init()
	generalbot.Start()

}
