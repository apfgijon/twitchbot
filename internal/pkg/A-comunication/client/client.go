package client

import "github.com/gempir/go-twitch-irc/v2"

type Communication struct {
	Client  *twitch.Client
	Channel string
	BotName string
}
