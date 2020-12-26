package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/apfgijon/cartones/internal/twitchbot"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var bot_name string
	var channel string
	var oauth string

	flag.StringVar(&bot_name, "b", "", "[required] Bot account name")
	flag.StringVar(&channel, "c", "", "[required] Channel to connect bot")
	flag.StringVar(&oauth, "o", "", "[required] Oauth key")

	flag.Parse()

	if bot_name == "" || channel == "" || oauth == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	twitchbot.Start(bot_name, channel, oauth)

	//riddle.StartClient(bot_name, channel, oauth)
}
