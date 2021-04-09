package commands

import (
	"github.com/apfgijon/cartones/internal/pkg/municipios"
	"github.com/gempir/go-twitch-irc/v2"
)

func (this *Commandsv1) checkArgCommands(message twitch.PrivateMessage, com string, args string) string {
	switch com {
	case "!municipio":
		resp := municipios.HablameSobre(args)
		return resp
	case "!quever":
		resp := municipios.QueVer(args)
		return resp
	case "!pokemon":
		return this.provider.GetPokemonRandomResponse(args)
	case "!ataques", "!moves":
		return this.provider.GetPokemonAtacksResponse(args)
	case "!tipo":
		return this.provider.GetPokemonTypesResponse(args)
	case "!evo":
		return this.provider.GetPokemonEvolutionResponse(args)
	case "!stats":
		return this.provider.GetPokemonStatsResponse(args)
	case "!tablatipos":
		return this.provider.GetPokemonTypeTableResponse(args)
	case "!capture", "!captura", "!rate":
		return this.provider.GetPokemonCaptureRateResponse(args)
	case "!covid":
		return this.provider.GetCovidStatsResponse(args, message.User.DisplayName)
	case "!botella":
		response := message.User.DisplayName + " tiró la botella y cayó en " + args
		return response
	default:
		return ""
	}
}
