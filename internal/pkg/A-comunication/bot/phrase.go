package bot

import (
	"time"

	"github.com/apfgijon/cartones/pkg/randomsay"
)

const timeToWait = time.Minute * 3

func (gn *Generalbot) sayRandomPhrase() {
	for {
		gn.com.Client.Say(gn.com.Channel, randomsay.GetRandomPhrase())
		time.Sleep(timeToWait)
	}
}

func (gn *Generalbot) sayRandomRefran() {
	for {
		gn.com.Client.Say(gn.com.Channel, randomsay.GetRandomRefran())
		time.Sleep(timeToWait)
	}
}
