package main

import (
	"flag"
	"fmt"
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
	var pokeGame string

	flag.StringVar(&botName, "b", "", "[required] Bot account name")
	flag.StringVar(&channel, "c", "", "[required] Channel to connect bot")
	flag.StringVar(&oauth, "o", "", "[required] Oauth key")
	flag.StringVar(&pokeGame, "pokeGame", "HG", "Pokemon Game Version")

	flag.Parse()

	if botName == "" || channel == "" || oauth == "" {
		fmt.Println("Wrong usage!")
		fmt.Println("bot -b <account> -c <channel> -o <oauth code>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	pokeapi.CacheSettings.CustomExpire = 1000000000000000

	bot := twitchbot.NewTwitchBot(botName, channel, oauth, pokeGame)
	bot.Bootstrap()
	//riddle.StartClient(botName, channel, oauth)
}
