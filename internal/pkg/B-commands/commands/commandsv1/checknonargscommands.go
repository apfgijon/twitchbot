package commands

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/mtslzr/pokeapi-go"
)

func (this *Commandsv1) checkNonArgCommands(message twitch.PrivateMessage) string {
	switch message.Message {
	case "!underlevel":
		num, _ := ioutil.ReadFile("under.txt")
		numI, _ := strconv.Atoi(string(num))
		numI++
		message := "El guiador dijo que está underlevel " + fmt.Sprint(numI) + " veces"
		ioutil.WriteFile("under.txt", []byte(fmt.Sprint(numI)), 0644)
		return message
	case "!carton":
		carton := cartongen.GenerateCarton()
		message := "Esti ye'l Bingu bot del mio canal, Equí ta'l to cartón " + message.User.DisplayName
		message = message + carton
		return message
	case "!pokemon":
		l, _ := pokeapi.Resource("pokemon", 0, 386)
		RandomPoke := l.Results[rand.Intn(len(l.Results))]
		response := message.User.DisplayName + " tiene la personalidad de " + RandomPoke.Name
		return response
	case "!botella":
		users, _ := this.users()
		user := rand.Intn(len(users))
		response := message.User.DisplayName + " tiró la botella y cayó en " + users[user]
		return response
	}
	return ""
}
