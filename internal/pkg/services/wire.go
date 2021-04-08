//+build wireinject

package services

import (
	"github.com/apfgijon/cartones/internal/pkg/A-comunication/bot"
	"github.com/apfgijon/cartones/internal/pkg/A-comunication/client"
	commands "github.com/apfgijon/cartones/internal/pkg/B-commands/commands/commandsv1"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
	"github.com/google/wire"
)

func InitializeBot(client.Communication) (bot.Bot, error) {
	wire.Build(pokemon.NewPokemonImpl, covid.NewCovidApi, bot.NewGeneralBot, commands.NewCommandImpl)
	return &bot.Generalbot{}, nil
}
