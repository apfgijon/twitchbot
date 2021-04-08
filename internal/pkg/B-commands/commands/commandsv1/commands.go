package commands

import (
	"strconv"
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/B-commands/commands"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
	"github.com/gempir/go-twitch-irc/v2"
)

const separator string = " ___________________________________________________ "

type Commandsv1 struct {
	poke  pokemon.PokeInfo
	covid covid.CovidInfo
	users func() ([]string, error)
}

func NewCommandImpl(p pokemon.PokeInfo, cov covid.CovidInfo) commands.Commands {
	return &Commandsv1{
		poke:  p,
		covid: cov,
	}
}

func (this *Commandsv1) Build(u func() ([]string, error)) {
	this.poke.Build()
	this.covid.Build()
	this.users = u
}

func (this *Commandsv1) CheckMessage(message twitch.PrivateMessage) string {

	ret := this.checkStaticCommands(message)

	if ret != "" {
		return ret
	}

	ret = this.checkNonArgCommands(message)

	if ret != "" {
		return ret
	}

	completeCommand := strings.Split(message.Message, " ")

	if len(completeCommand) > 1 {
		com := completeCommand[0]
		args := strings.Join(completeCommand[1:], " ")

		com = strings.ToLower(com)

		ret = this.checkArgCommands(message, com, args)

		if ret != "" {
			return ret
		}
	}

	return ""
}

func (this *Commandsv1) covidStats(message twitch.PrivateMessage, args string) string {
	transalatedresp := this.covid.FormatName(args)
	if transalatedresp != "" {
		casos, muertos := this.covid.GetCovidCasesForProvince(transalatedresp)
		formattedMessage := "Casos de covid de güei d'" + args + separator +
			"Casos novos güei: " + strconv.Itoa(casos) + separator +
			"Mortos güei: " + strconv.Itoa(muertos)
		if casos == 0 {
			formattedMessage = "Vaya, parece que no tengo datos hoy @" + message.User.DisplayName + " :("
		}
		return formattedMessage
	}
	args = strings.ToLower(args)
	if args == "españa" {
		casos, muertos := this.covid.GetCovidCasesSpain()

		formattedMessage := "Casos de covid de güei d'" + args + separator +
			"Casos novos güei: " + strconv.Itoa(casos) + separator +
			"Mortos güei: " + strconv.Itoa(muertos)
		if casos == 0 {
			formattedMessage = "Vaya, parece que no tengo datos hoy @" + message.User.DisplayName + " :("
		}
		return formattedMessage

	}
	formattedMessage := "Nun sei " + message.User.DisplayName + ", abondo que poño casos d'españa"
	return formattedMessage
}

func (this *Commandsv1) ataques(args string, game string) string {
	return this.poke.PokeMovesFormatted(args, game)
}

func (this *Commandsv1) evolution(args string) string {
	return this.poke.PokeEvos(args)
}

func (this *Commandsv1) captureRate(args string) string {
	return "Ratio de captura de " + args + ": " + strconv.Itoa(this.poke.CaptureRate(args))
}

func (this *Commandsv1) tipos(args string) string {
	tipo := this.poke.Types(args)
	if tipo == "" {
		return ""
	}
	return args + " es tipo " + tipo
}

func (this *Commandsv1) stats(args string) string {
	return this.poke.Stats(args)
}

func (this *Commandsv1) tablatiposHacia(args string) string {
	message := this.poke.TypeTableTo(args)

	if message == "" {
		return ""
	}

	return message
}

func (this *Commandsv1) tablatiposDe(args string) string {
	message := this.poke.TypeTableFrom(args)

	if message == "" {
		return ""
	}

	return message
}

func (this *Commandsv1) tablatipos(args string) string {
	message := this.poke.TypeTable(args)

	if message == "" {
		return ""
	}

	return message
}
