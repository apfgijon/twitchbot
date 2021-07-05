package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	case "!peso":
		return this.provider.GetPokemonPesoResponse(args)
	case "!tabla", "!tablatipos":
		return this.provider.GetPokemTable(args)
	case "!capture", "!captura", "!rate":
		return this.provider.GetPokemonCaptureRateResponse(args)
	case "!covid":
		return this.provider.GetCovidStatsResponse(args, message.User.DisplayName)
	case "!botella":
		response := message.User.DisplayName + " tiró la botella y cayó en " + args
		return response
	case "!pp":
		return this.provider.GetPPResponse(args)
	case "!coin":
		r, err := http.Get("https://api.coincap.io/v2/assets?search=" + args)
		if err != nil {
			return ""
		}
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return ""
		}
		var c Coin
		json.Unmarshal(body, &c)

		if len(c.Data) == 0 {
			r, err := http.Get("https://api.coincap.io/v2/assets?ids=" + args)

			if err != nil {
				return ""
			}
			defer r.Body.Close()
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return ""
			}
			json.Unmarshal(body, &c)
			if len(c.Data) == 0 {
				return ""
			}
		}

		return "La " + args + " ta a " + c.Data[0].Priceusd + " USD"
	default:
		return ""
	}
}
