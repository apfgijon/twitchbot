package bot

import (
	"math/rand"
	"time"

	"github.com/apfgijon/cartones/pkg/randomsay"
)

const timeToWait = time.Minute * 3

func (gn *Generalbot) sayRandomPhrase() {
	for {
		gn.com.Client.Say(gn.com.Channel, randomsay.GetRandomPhrase())

		ran := rand.Intn(1)

		time.Sleep(timeToWait * time.Duration(ran+1))
	}
}

func (gn *Generalbot) sayRandomRefran() {
	for {
		gn.com.Client.Say(gn.com.Channel, randomsay.GetRandomRefran())

		time.Sleep(timeToWait)
	}
}
