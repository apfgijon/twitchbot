package prov

import (
	"strconv"
	"strings"

	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
)

const separator string = " ___________________________________________________ "

type MessageProviderv1 struct {
	poke  pokemon.PokeInfo
	covid covid.CovidInfo
	car   cartongen.Carton
}

func NewMessageProoviderv1(p pokemon.PokeInfo, cov covid.CovidInfo, c cartongen.Carton) MessageProvider {
	return &MessageProviderv1{
		poke:  p,
		covid: cov,
		car:   c,
	}
}

func (this *MessageProviderv1) Build() {
	this.poke.Build()
	this.covid.Build()
}

func (this *MessageProviderv1) GetBingoCartonResponse() string {
	return "Esti ye'l Bingu bot del mio canal, Equí ta'l to cartón " + this.car.GenerateCarton()
}

func (this *MessageProviderv1) GetPokemonRandomResponse(user string) string {
	return user + " tiene la personalidad de " + this.poke.PokeRandom()
}

func (this *MessageProviderv1) GetPokemonAtacksResponse(poke string) string {
	return this.poke.PokeMovesFormatted(poke)
}

func (this *MessageProviderv1) GetPokemonEvolutionResponse(poke string) string {
	return this.poke.PokeEvos(poke)
}

func (this *MessageProviderv1) GetPokemonTypesResponse(poke string) string {
	tipo := this.poke.Types(poke)
	if tipo == "" {
		return ""
	}
	return poke + " es tipo " + tipo
}

func (this *MessageProviderv1) GetPokemonCaptureRateResponse(poke string) string {
	return "Ratio de captura de " + poke + ": " + strconv.Itoa(this.poke.CaptureRate(poke))
}

func (this *MessageProviderv1) GetPokemonStatsResponse(poke string) string {
	return this.poke.Stats(poke)
}

func (this *MessageProviderv1) GetPokemonTypeTableResponse(typ string) string {
	message := this.poke.TypeTable(typ)

	if message == "" {
		return ""
	}

	return message
}

func (this *MessageProviderv1) GetCovidStatsResponse(site string, user string) string {
	transalatedresp := this.covid.FormatName(site)
	if transalatedresp != "" {
		casos, muertos := this.covid.GetCovidCasesForProvince(transalatedresp)
		formattedMessage := "Casos de covid de güei d'" + site + separator +
			"Casos novos güei: " + strconv.Itoa(casos) + separator +
			"Mortos güei: " + strconv.Itoa(muertos)
		if casos == 0 {
			formattedMessage = "Vaya, parece que no tengo datos hoy @" + user + " :("
		}
		return formattedMessage
	}
	site = strings.ToLower(site)
	if site == "españa" {
		casos, muertos := this.covid.GetCovidCasesSpain()

		formattedMessage := "Casos de covid de güei d'" + site + separator +
			"Casos novos güei: " + strconv.Itoa(casos) + separator +
			"Mortos güei: " + strconv.Itoa(muertos)
		if casos == 0 {
			formattedMessage = "Vaya, parece que no tengo datos hoy @" + user + " :("
		}
		return formattedMessage

	}
	formattedMessage := "Nun sei " + user + ", abondo que poño casos d'españa"
	return formattedMessage
}
