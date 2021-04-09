package commands

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"

	"github.com/gempir/go-twitch-irc/v2"
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
		return this.provider.GetBingoCartonResponse()
	case "!pokemon":
		return this.provider.GetPokemonRandomResponse(message.User.DisplayName)
	case "!botella":
		users, _ := this.users()
		user := rand.Intn(len(users))
		response := message.User.DisplayName + " tiró la botella y cayó en " + users[user]
		return response
	}
	return ""
}
