package commands

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gempir/go-twitch-irc/v2"
)

type Coin struct {
	Data []struct {
		Priceusd string `json:"priceUsd"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

func (this *Commandsv1) checkNonArgCommands(message twitch.PrivateMessage) string {
	switch message.Message {
	case "!+iq", "!masiq":

		r := rand.Intn(10)

		iqRaw, _ := ioutil.ReadFile("iq")

		iq, _ := strconv.Atoi(string(iqRaw))
		luck := ""
		sum := 0
		if r < 5 {
			luck = "Bueeeno habló a la minita que le gusta"
			iq = iq + 1
			sum = 1
		} else if r < 8 {
			sum = 5
			luck = "Bastante bien aprobó el examen del salón"
			iq = iq + 5
		} else if r < 9 {
			sum = 10
			luck = "Joder aprendió a atarse los cordones con 8 años"
			iq = iq + 10
		} else {
			sum = 20
			luck = "Encontró la funete del iq"
			iq = iq + 20
		}

		message := "El guiador consumió su tremenda energía para ser mucho más inteligente!." + luck + " Ahora el guiador tiene " + strconv.Itoa(iq) + " IQ sumando " + strconv.Itoa(sum)

		ioutil.WriteFile("iq", []byte(strconv.Itoa(iq)), 0644)

		return message
	case "!-iq", "!menosiq":

		r := rand.Intn(10)

		iqRaw, _ := ioutil.ReadFile("iq")

		iq, _ := strconv.Atoi(string(iqRaw))

		sum := 0
		luck := ""
		if r < 5 {
			sum = 1
			luck = "No se cago tanto"
			iq = iq - 1
		} else if r < 8 {
			sum = 5
			luck = "Cagose un poco"
			iq = iq - 5
		} else if r < 9 {
			sum = 10
			luck = "Jaja se mió"
			iq = iq - 10
		} else {
			sum = 20
			luck = "Ta esfoirao"
			iq = iq - 20
		}

		if iq < 0 {
			return "Lo siento no se puede tener menos de 0 iq porque es Años/Años y no se pueden tener años negativos"
		}

		message := "El guiador se hizo caca en los pantalones :( ." + luck + " Ahora el guiador tiene " + strconv.Itoa(iq) + " IQ restando " + strconv.Itoa(sum)

		ioutil.WriteFile("iq", []byte(strconv.Itoa(iq)), 0644)

		return message
	case "!iq":

		iqRaw, _ := ioutil.ReadFile("iq")

		iq, _ := strconv.Atoi(string(iqRaw))

		message := "El guiador tiene " + strconv.Itoa(iq) + " IQ"

		return message
	case "!carton":
		return this.provider.GetBingoCartonResponse()
	case "!pokemon":
		return this.provider.GetPokemonRandomResponse(message.User.DisplayName)
	case "!botella":
		return this.provider.GetBotellaResponse(this.users, message.User.DisplayName)
	case "!dogecoin":
		r, _ := http.Get("https://api.coincap.io/v2/assets?search=dogecoin")
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		var c Coin
		json.Unmarshal(body, &c)

		dcPrevRaw, _ := ioutil.ReadFile("dogecoin")

		dcPrev, _ := strconv.ParseFloat(string(dcPrevRaw), 32)

		ioutil.WriteFile("dogecoin", []byte(c.Data[0].Priceusd), 0644)

		dcNow, _ := strconv.ParseFloat(c.Data[0].Priceusd, 32)

		percentaje := ((dcNow / dcPrev) * 100) - 100

		if dcNow > dcPrev {
			return "El DogeCoin ta a " + c.Data[0].Priceusd + " USD, subió un " + strconv.FormatFloat(percentaje, 'f', 2, 64) + "% desde la última vez CorgiDerp"
		}
		return "El DogeCoin ta a " + c.Data[0].Priceusd + " USD, bajó un " + strconv.FormatFloat(percentaje, 'f', 2, 64) + "% desde la última vez CorgiDerp"

	}
	return ""
}
