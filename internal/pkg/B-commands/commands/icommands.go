package commands

import "github.com/gempir/go-twitch-irc/v2"

type Commands interface {
	CheckMessage(message twitch.PrivateMessage) string
	SetUsersConnectedProvider(users func() ([]string, error))
}
