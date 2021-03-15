package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"github.com/apfgijon/cartones/internal/twitchbot"
	"github.com/mtslzr/pokeapi-go"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var botName string
	var channel string
	var oauth string

	flag.StringVar(&botName, "b", "", "[required] Bot account name")
	flag.StringVar(&channel, "c", "", "[required] Channel to connect bot")
	flag.StringVar(&oauth, "o", "", "[required] Oauth key")

	flag.Parse()

	if botName == "" || channel == "" || oauth == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	pokeapi.CacheSettings.CustomExpire = 1000000000000000

	twitchbot.Start(botName, channel, oauth)
	//riddle.StartClient(botName, channel, oauth)
}
