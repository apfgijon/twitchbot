package generalbot

import (
	"strconv"
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/municipios"
	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/gempir/go-twitch-irc/v2"
)

func (gn *Generalbot) checkCommands(message twitch.PrivateMessage) bool {
	switch message.Message {
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
	case "!javi":
		message := "Nah un putu tryhard de la de dios"
		gn.Com.Client.Say(gn.Com.Channel, message)
		break
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
		case "!covid":
			transalatedresp := covid.Translate(args)
			if transalatedresp != "" {
				resp := covid.GetCovidCasesForProvince(transalatedresp)
				formattedMessage := "Casos de covid de güei d'" + args + " ___________________________________________________ " +
					"Casos novos güei: " + strconv.Itoa(resp.ConfirmedDiff) + " ___________________________________________________ " +
					"Mortos güei: " + strconv.Itoa(resp.DeathsDiff)
				gn.Com.Client.Say(gn.Com.Channel, formattedMessage)
				return true
			}
			args = strings.ToLower(args)
			if args == "españa" {
				casos, muertos := covid.GetCovidCasesSpain()
				formattedMessage := "Casos de covid de güei d'" + args + " ___________________________________________________ " +
					"Casos novos güei: " + strconv.Itoa(casos) + " ___________________________________________________ " +
					"Mortos güei: " + strconv.Itoa(muertos)
				gn.Com.Client.Say(gn.Com.Channel, formattedMessage)
				return true

			}
			break
		default:
			return false
		}

	}

	return false
}
