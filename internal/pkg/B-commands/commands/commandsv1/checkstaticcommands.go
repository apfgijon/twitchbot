package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

func (this *Commandsv1) checkStaticCommands(message twitch.PrivateMessage) string {
	commandsRaw, err := ioutil.ReadFile("commands.json")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	commands := make(map[string][]string)

	json.Unmarshal(commandsRaw, &commands)

	messageComm := strings.Split(message.Message, " ")

	for i, v := range commands {

		if strings.ToLower(i) == strings.ToLower(messageComm[0]) {
			var ran int
			var err error
			if len(messageComm) > 1 {
				ran, err = strconv.Atoi(messageComm[1])
				ran--

			}
			if err != nil || ran < 0 || ran >= len(v) {
				ran = rand.Intn(len(v))
			}

			send := strings.ReplaceAll(v[ran], "$Name", message.User.DisplayName)
			return send
		}
	}

	return ""
}
