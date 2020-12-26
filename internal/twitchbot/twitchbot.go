package twitchbot

import (
	"github.com/apfgijon/cartones/internal/pkg/client"
	"github.com/apfgijon/cartones/internal/pkg/generalbot"
	"github.com/gempir/go-twitch-irc/v2"
)

func Start(bot_name string, channel string, oauth string) {

	client := client.Communication{
		Client:  twitch.NewClient(bot_name, oauth),
		Channel: channel,
		BotName: bot_name,
	}

	generalbot := generalbot.Generalbot{
		Com: client,
	}

	generalbot.Init()
	generalbot.Start()

}
