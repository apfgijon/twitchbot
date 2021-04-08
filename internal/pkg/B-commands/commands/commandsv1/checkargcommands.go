package commands

import (
	"math/rand"

	"github.com/apfgijon/cartones/internal/pkg/municipios"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/mtslzr/pokeapi-go"
)

func (this *Commandsv1) checkArgCommands(message twitch.PrivateMessage, com string, args string) string {
	switch com {
	case "!municipio":

		resp := municipios.HablameSobre(args)
		return resp
	case "!pokemon":
		l, _ := pokeapi.Resource("pokemon", 0, 386)
		RandomPoke := l.Results[rand.Intn(len(l.Results))]
		response := args + " tiene la personalidad de " + RandomPoke.Name
		return response
	case "!quever":
		resp := municipios.QueVer(args)
		return resp
	case "!ataques", "!moves":
		return this.ataques(args, "heartgold-soulsilver")
	case "!tipo":
		return this.tipos(args)
	case "!evo":
		return this.evolution(args)
	case "!stats":
		return this.stats(args)
	case "!tiposatacar", "!tipoatacar", "!efectivo":
		return this.tablatiposHacia(args)
	case "!tiposrecibir", "!tiporecibir", "!resiste", "!weak":
		return this.tablatiposDe(args)
	case "!tablatipos":
		return this.tablatipos(args)
	case "!capture", "!captura", "!rate":
		return this.captureRate(args)
	case "!covid":
		return this.covidStats(message, args)
	case "!botella":
		response := message.User.DisplayName + " tiró la botella y cayó en " + args
		return response
	default:
		return ""
	}
}
