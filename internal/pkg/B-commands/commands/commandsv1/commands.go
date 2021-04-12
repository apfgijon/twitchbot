package commands

import (
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/B-commands/commands"
	"github.com/apfgijon/cartones/internal/pkg/C-style/prov"
	"github.com/gempir/go-twitch-irc/v2"
)

const separator string = " ___________________________________________________ "

type Commandsv1 struct {
	provider prov.MessageProvider
	users    func() ([]string, error)
}

func NewCommandImpl(p prov.MessageProvider) (commands.Commands, error) {
	this := &Commandsv1{
		provider: p,
	}

	this.users = func() ([]string, error) {
		return []string{}, nil
	}

	return this, nil
}

func (this *Commandsv1) SetUsersConnectedProvider(u func() ([]string, error)) {
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
