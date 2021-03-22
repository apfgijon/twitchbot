package generalbot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/municipios"
	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/mtslzr/pokeapi-go"
)

const separator string = " ___________________________________________________ "

func (gn *Generalbot) checkCommands(message twitch.PrivateMessage) bool {

	ret := gn.checkStaticCommands(message)

	if ret == true {
		return ret
	}

	switch message.Message {
	case "!carton":
		carton := cartongen.GenerateCarton()
		message := "Esti ye'l Bingu bot del mio canal, Equí ta'l to cartón " + message.User.DisplayName + "                                                  "
		message = message + carton
		gn.com.Client.Say(gn.com.Channel, message)
		break
	case "!albur", "!hot", "!colorled":
		message := "Que faes usando Nightbot tando yo equí " + message.User.DisplayName + "? Yo de verdá nun pescancio res"
		gn.com.Client.Say(gn.com.Channel, message)
		break
	case "!pokemon":
		l, _ := pokeapi.Resource("pokemon", 0, 386)
		RandomPoke := l.Results[rand.Intn(len(l.Results))]
		response := message.User.DisplayName + " tiene la personalidad de " + RandomPoke.Name
		gn.com.Client.Say(gn.com.Channel, response)
		break
	}

	completeCommand := strings.Split(message.Message, " ")

	if len(completeCommand) > 1 {
		com := completeCommand[0]
		args := strings.Join(completeCommand[1:], " ")

		switch com {
		case "!municipio":

			resp := municipios.HablameSobre(args)
			gn.com.Client.Say(gn.com.Channel, resp)
			break
		case "!pokemon":
			l, _ := pokeapi.Resource("pokemon", 0, 386)
			RandomPoke := l.Results[rand.Intn(len(l.Results))]
			response := args + " tiene la personalidad de " + RandomPoke.Name
			gn.com.Client.Say(gn.com.Channel, response)
			break

		case "!quever":

			resp := municipios.QueVer(args)
			gn.com.Client.Say(gn.com.Channel, resp)
			break
		case "!ataques", "!moves":
			go gn.ataques(args, "heartgold-soulsilver")
			break
		case "!tipo":
			go gn.tipos(args)
			break
		case "!evo":
			go gn.evolution(args)
			break
		case "!stats":
			go gn.stats(args)
			break
		case "!capture", "!captura":
			go gn.captureRate(args)
			break
		case "!covid":
			transalatedresp := gn.covid.FormatName(args)
			if transalatedresp != "" {
				casos, muertos := gn.covid.GetCovidCasesForProvince(transalatedresp)
				formattedMessage := "Casos de covid de güei d'" + args + separator +
					"Casos novos güei: " + strconv.Itoa(casos) + separator +
					"Mortos güei: " + strconv.Itoa(muertos)
				if casos == 0 {
					formattedMessage = "Vaya, parece que no tengo datos hoy @" + message.User.DisplayName + " :("
				}
				gn.com.Client.Say(gn.com.Channel, formattedMessage)
				return true
			}
			args = strings.ToLower(args)
			if args == "españa" {
				casos, muertos := gn.covid.GetCovidCasesSpain()

				formattedMessage := "Casos de covid de güei d'" + args + separator +
					"Casos novos güei: " + strconv.Itoa(casos) + separator +
					"Mortos güei: " + strconv.Itoa(muertos)
				if casos == 0 {
					formattedMessage = "Vaya, parece que no tengo datos hoy @" + message.User.DisplayName + " :("
				}
				gn.com.Client.Say(gn.com.Channel, formattedMessage)
				return true

			}
			formattedMessage := "Nun sei " + message.User.DisplayName + ", abondo que poño casos d'españa"
			gn.com.Client.Say(gn.com.Channel, formattedMessage)
			break
		default:
			return false
		}

	}

	return false
}

func (gn *Generalbot) checkStaticCommands(message twitch.PrivateMessage) bool {
	commandsRaw, err := ioutil.ReadFile("commands.json")

	if err != nil {
		fmt.Println(err)
		return false
	}

	commands := make(map[string]string)

	json.Unmarshal(commandsRaw, &commands)

	for i, v := range commands {
		if i == message.Message {
			gn.com.Client.Say(gn.com.Channel, v)
			return true
		}
	}

	return false
}

func (gn *Generalbot) ataques(args string, game string) {
	gn.com.Client.Say(gn.com.Channel, gn.poke.PokeMovesFormatted(args, game))
}

func (gn *Generalbot) evolution(args string) {
	gn.com.Client.Say(gn.com.Channel, gn.poke.PokeEvos(args))
}

func (gn *Generalbot) captureRate(args string) {
	gn.com.Client.Say(gn.com.Channel, "Ratio de captura de "+args+": "+strconv.Itoa(gn.poke.CaptureRate(args)))
}

func (gn *Generalbot) tipos(args string) {
	gn.com.Client.Say(gn.com.Channel, gn.poke.Types(args))
}

func (gn *Generalbot) stats(args string) {
	gn.com.Client.Say(gn.com.Channel, gn.poke.Stats(args))
}
