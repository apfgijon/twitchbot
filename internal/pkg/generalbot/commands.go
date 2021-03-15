package generalbot

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/municipios"
	"github.com/apfgijon/cartones/internal/pkg/pokemon"
	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/mtslzr/pokeapi-go"
)

const separator string = " ___________________________________________________ "

func (gn *Generalbot) checkCommands(message twitch.PrivateMessage) bool {
	switch message.Message {
	case "gg":
		message := "!yellow"
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!carton":
		carton := cartongen.GenerateCarton()
		message := "Esti ye'l Bingu bot del mio canal, Equí ta'l to cartón " + message.User.DisplayName + "                                                  "
		message = message + carton
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!albur", "!hot", "!colorled":
		message := "Que faes usando Nightbot tando yo equí " + message.User.DisplayName + "? Yo de verdá nun pescancio res"
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!gonzalo":
		message := "shhhh nun fales de \"E LOGO\""
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!pokemon":
		l, _ := pokeapi.Resource("pokemon", 0, 386)
		RandomPoke := l.Results[rand.Intn(len(l.Results))]
		response := message.User.DisplayName + " tiene la personalidad de " + RandomPoke.Name
		gn.Com.Client.Say(gn.Com.Channel, response)
		break
	case "!muertos":
		message := "Rython(Haz_A), Zigzagoon(JavvyoYT), Tyrogue(Yajuli), Shelgon(Mr Socone), Crabby(lijo96), Relicanth(ZonnyoYT), Dodrio(lilwest), Golem(chinchypan), Charizard(Gonzalo), Kabutops(Mia), Crobat(bicicletis), Ampharos(wrysp), Cloyster(hetwan), Mewtwo(KAY) y Sceptile(Alick) :("
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!javi":
		message := "Nah un putu tryhard de la de dios"
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!social":
		message := "NO SOY"
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
	case "!skill":
		message := "https://clips.twitch.tv/BloodyColdbloodedShrewNotATK"
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
		// case "!alexa":
		// 	number := alexa.GetNumber()
		// 	message := "Caloto ha discutido con alexa " + strconv.Itoa(number) + " veces en stream"
		// 	gn.Com.Client.Say(gn.Com.Channel, message)
		// 	break
	}

	completeCommand := strings.Split(message.Message, " ")

	if len(completeCommand) > 1 {
		com := completeCommand[0]
		args := strings.Join(completeCommand[1:], " ")

		switch com {
		case "!municipio":

			resp := municipios.HablameSobre(args)
			gn.Com.Client.Say(gn.Com.Channel, resp)
			break

		case "!quever":

			resp := municipios.QueVer(args)
			gn.Com.Client.Say(gn.Com.Channel, resp)
			break
		case "!ataques":
			go ataques(gn, args)
			break
		case "!comoes":
			p, _ := pokeapi.Pokemon(strings.ToLower(args))
			resp := p.Sprites.FrontDefault
			gn.Com.Client.Say(gn.Com.Channel, resp)
			break
		case "!comoesshiny":
			p, _ := pokeapi.Pokemon(strings.ToLower(args))
			resp := p.Sprites.FrontShiny
			gn.Com.Client.Say(gn.Com.Channel, resp)
			break
		case "!tipo":
			go tipos(gn, args)
			break
		case "!covid":
			// if time.Now().Weekday() == 0 || time.Now().Weekday() == 5 || time.Now().Weekday() == 6 {
			// 	gn.Com.Client.Say(gn.Com.Channel, message.User.DisplayName+", nun hai datos güei")
			// 	return true
			// }
			transalatedresp := covid.Translate(args)
			if transalatedresp != "" {
				resp := covid.GetCovidCasesForProvince(transalatedresp)
				formattedMessage := "Casos de covid de güei d'" + args + separator +
					"Casos novos güei: " + strconv.Itoa(resp.ConfirmedDiff) + separator +
					"Mortos güei: " + strconv.Itoa(resp.DeathsDiff)
				gn.Com.Client.Say(gn.Com.Channel, formattedMessage)
				return true
			}
			args = strings.ToLower(args)
			if args == "españa" {
				casos, muertos := covid.GetCovidCasesSpain()
				formattedMessage := "Casos de covid de güei d'" + args + separator +
					"Casos novos güei: " + strconv.Itoa(casos) + separator +
					"Mortos güei: " + strconv.Itoa(muertos)
				gn.Com.Client.Say(gn.Com.Channel, formattedMessage)
				return true

			}
			formattedMessage := "Nun sei " + message.User.DisplayName + ", abondo que poño casos d'españa"
			gn.Com.Client.Say(gn.Com.Channel, formattedMessage)
			break
		default:
			return false
		}

	}

	return false
}

func ataques(gn *Generalbot, args string) {
	p, _ := pokeapi.Pokemon(strings.ToLower(args))

	moves := pokemon.PokeMoves(p, "emerald")
	gn.Com.Client.Say(gn.Com.Channel, moves)
}

func tipos(gn *Generalbot, args string) {
	p, _ := pokeapi.Pokemon(strings.ToLower(args))
	if p.Name == "" {
		return
	}
	typo := p.Types
	message := p.Name + " es tipo: "
	for _, v := range typo {
		TypeName := v.Type.Name
		// Tipo, _ := pokeapi.Type(TypeName)
		// for _, t := range Tipo.Names {
		// 	if t.Language.Name == "es" {
		// 		TypeName = t.Name
		// 	}
		// }

		message += TypeName + " "
	}
	gn.Com.Client.Say(gn.Com.Channel, message)
}
