package bot

import (
	"fmt"
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/A-comunication/client"
	"github.com/apfgijon/cartones/internal/pkg/B-commands/commands"
	"github.com/apfgijon/cartones/pkg/randomsay"
	"github.com/gempir/go-twitch-irc/v2"
)

type Bot interface {
	Start()
	Stop()
}

type Generalbot struct {
	com      client.Communication
	commands commands.Commands
}

func NewGeneralBot(commands commands.Commands, comu client.Communication) (Bot, error) {
	return &Generalbot{
		com:      comu,
		commands: commands,
	}, nil
}

func (this *Generalbot) Start() {
	this.commands.SetUsersConnectedProvider(this.userList)

	this.com.Client.OnPrivateMessage(this.onMessage)

	this.com.Client.Join(this.com.Channel)
	go this.sayRandomPhrase()
	// go this.sayRandomRefran()
	err := this.com.Client.Connect()
	if err != nil {
		fmt.Println("Desconectado")
	}
}
func (this *Generalbot) Stop() {
	this.com.Client.Disconnect()
}

func (this *Generalbot) onMessage(message twitch.PrivateMessage) {
	if message.User.DisplayName != "duendrixx" {

		go this.checkCommands(message)
	}

	if string(message.Message[0]) != "!" && string(message.Message[0]) != "@" && message.User.DisplayName != "Nightbot" && !strings.Contains(strings.ToLower(message.Message), "zonnyo") {
		go randomsay.SetPhrase(message.Message)

	}
}

func (this *Generalbot) userList() ([]string, error) {
	return this.com.Client.Userlist(this.com.Channel)
}

func (this *Generalbot) checkCommands(message twitch.PrivateMessage) {
	m := this.commands.CheckMessage(message)

	this.com.Client.Say(this.com.Channel, m)
}
