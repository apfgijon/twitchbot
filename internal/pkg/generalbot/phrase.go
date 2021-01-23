package generalbot

import (
	"time"

	"github.com/apfgijon/cartones/pkg/randomsay"
)

const timeToWait = time.Minute * 5

func (gn *Generalbot) sayRandomPhrase() {
	for {
		gn.Com.Client.Say(gn.Com.Channel, randomsay.GetRandomPhrase())
		time.Sleep(timeToWait)
	}
}

func (gn *Generalbot) sayRandomRefran() {
	for {
		gn.Com.Client.Say(gn.Com.Channel, randomsay.GetRandomRefran())
		time.Sleep(timeToWait)
	}
}
