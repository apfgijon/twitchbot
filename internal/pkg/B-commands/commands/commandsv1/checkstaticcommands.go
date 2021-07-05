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

type Comm struct {
	Text    []string `json:"text"`
	Counter int      `json:"counter"`
}

func (this *Commandsv1) checkStaticCommands(message twitch.PrivateMessage) string {
	commandsRaw, err := ioutil.ReadFile("commands.json")

	if err != nil {
		fmt.Println(err)
		return ""
	}
	send := ""
	commands := make(map[string]Comm)

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
			if err != nil || ran < 0 || ran >= len(v.Text) || len(messageComm) == 1 {

				ran = rand.Intn(len(v.Text))
			}

			send = strings.ReplaceAll(v.Text[ran], "$Name", message.User.DisplayName)
			send = strings.ReplaceAll(send, "$Counter", strconv.Itoa(v.Counter))

			commands[i] = Comm{
				Text:    commands[i].Text,
				Counter: commands[i].Counter + 1,
			}
		}
	}
	b, _ := json.MarshalIndent(commands, "", " ")

	ioutil.WriteFile("commands.json", b, 0644)
	return send
}
