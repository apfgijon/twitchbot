package prov

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/D-filesystem/filesystem"
	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
)

const separator string = " ___________________________________________________ "

type MessageProviderv1 struct {
	poke  pokemon.PokeInfo
	covid covid.CovidInfo
	car   cartongen.Carton
	fs    filesystem.FileProvider
}

func NewMessageProoviderv1(p pokemon.PokeInfo, cov covid.CovidInfo, c cartongen.Carton, f filesystem.FileProvider) (MessageProvider, error) {
	this := &MessageProviderv1{
		poke:  p,
		covid: cov,
		car:   c,
		fs:    f,
	}
	return this, nil
}

func (this *MessageProviderv1) Build() {

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
	rate := this.poke.CaptureRate(poke)
	if rate == 0 {
		return ""
	}
	return "Ratio de captura de " + poke + ": " + strconv.Itoa(rate)
}

func (this *MessageProviderv1) GetPokemonStatsResponse(poke string) string {
	return this.poke.Stats(poke)
}

func (this *MessageProviderv1) GetPokemonPesoResponse(poke string) string {
	return this.poke.Peso(poke)
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

func (this *MessageProviderv1) GetBotellaResponse(getusers func() ([]string, error), u string) string {
	users, _ := getusers()
	user := rand.Intn(len(users))
	response := u + " tiró la botella y cayó en " + users[user]
	return response
}

func (this *MessageProviderv1) GetPokemTable(poke string) string {
	esPokemon, TypesFrom, TypesTo := this.poke.TypeTablePokemon(poke)

	if esPokemon == 0 {
		return ""
	}
	var ret string
	if esPokemon == 1 {

		ret = "Pokemon: " + poke
	} else {
		ret = "Tipo: " + poke
	}

	ya := true

	for i, v := range TypesFrom {
		if v > 0 {
			if ya {
				ret += " | LU JODE: "
				ya = false
			}
			ret += i
			if v == 2 {
				ret += "(x4)"
			}
			ret += " "
		}
	}
	ya = true
	for i, v := range TypesFrom {
		if v < 0 {
			if ya {
				ret += " | RESISTE: "
				ya = false
			}
			ret += i
			if v == -2 {
				ret += "(x1/4)"
			}
			ret += " "
		}
	}
	ya = true
	for i, v := range TypesFrom {
		if v < -5 {
			if ya {
				ret += " | NO LO AFECTA: "
				ya = false
			}
			ret += i + " "
		}
	}
	if esPokemon == 2 {
		ya = true
		for i, v := range TypesTo {
			if v > 0 {
				if ya {
					ret += " | EFECTIVO CONTRA: "
					ya = false
				}
				ret += i + " "
			}
		}
		ya = true
		for i, v := range TypesTo {
			if v < 0 {
				if ya {
					ret += " | PUTA MIERDA CONTRA: "
					ya = false
				}
				ret += i + " "
			}
		}
		ya = true
		for i, v := range TypesTo {
			if v < -5 {
				if ya {
					ret += " | NO EFECTIVO CONTRA: "
					ya = false
				}
				ret += i + " "
			}
		}
	}

	return ret
}

func (this *MessageProviderv1) GetPPResponse(a string) string {
	return "El ataque " + a + " tiene " + strconv.Itoa(this.poke.PP(a)) + " PPs"
}
