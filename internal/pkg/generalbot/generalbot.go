package generalbot

import (
	"strings"

	"github.com/apfgijon/cartones/internal/pkg/client"
	"github.com/apfgijon/cartones/pkg/randomsay"
	"github.com/gempir/go-twitch-irc/v2"
)

type Generalbot struct {
	Com          client.Communication
	JavvyoYTesta bool
	HannyaYTesta bool
	Trollchuesta bool
	Haz_Aesta    bool
	chisseiesta  bool
	zaraaify     bool
	mariana      bool
	miamaguila   bool
}

func (gn *Generalbot) Init() {
	gn.JavvyoYTesta = false
	gn.HannyaYTesta = false
	gn.Trollchuesta = false
	gn.Haz_Aesta = false
	gn.chisseiesta = false
	gn.zaraaify = false
	gn.mariana = false
	gn.miamaguila = false

}

func (gn *Generalbot) Start() {
	gn.Com.Client.OnPrivateMessage(gn.onMessage)

	gn.Com.Client.Join(gn.Com.Channel)
	go gn.sayRandomPhrase()
	err := gn.Com.Client.Connect()
	if err != nil {
		panic(err)
	}
}

func (gn *Generalbot) onMessage(message twitch.PrivateMessage) {
	gn.salute(message)
	if !gn.checkCommands(message) {
		if strings.Contains(strings.ToUpper(message.Message), strings.ToUpper(gn.Com.BotName)) {
			message := "Que me dices " + message.User.DisplayName + "? nun ves que soy un bot? Amás nun pescancio castel.lán."
			gn.Com.Client.Say(gn.Com.Channel, message)
		}
		randomsay.SetPhrase(message.Message)
	}
}