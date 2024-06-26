package riddle

import (
	"os"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
)

func StartClient(botName string, channel string, oauth string) {
	client := twitch.NewClient(botName, oauth)

	riddle(client, channel)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		mensaje := strings.ToLower(message.Message)

		if strings.Contains(mensaje, "rocío") {
			client.Say(channel, message.User.DisplayName+" Enhorabuena solo si no eres javi (Javi dejas de tryhardear ya?)")
			time.Sleep(10)
			os.Exit(1)
		}

	})

	client.Join(channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func riddle(client *twitch.Client, channel string) {
	message := ""

	client.Say(channel, message)
}
