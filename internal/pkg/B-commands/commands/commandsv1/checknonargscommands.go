package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
)

func (this *Commandsv1) checkNonArgCommands(message twitch.PrivateMessage) string {
	switch message.Message {
	case "!underlevel":
		return this.provider.GetUnderLevelResponse("!underlevel")
	case "!carton":
		return this.provider.GetBingoCartonResponse()
	case "!pokemon":
		return this.provider.GetPokemonRandomResponse(message.User.DisplayName)
	case "!botella":
		return this.provider.GetBotellaResponse(this.users, message.User.DisplayName)
	}
	return ""
}
