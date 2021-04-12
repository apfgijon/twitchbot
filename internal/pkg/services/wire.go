//+build wireinject

package services

import (
	"github.com/apfgijon/cartones/internal/pkg/A-comunication/bot"
	"github.com/apfgijon/cartones/internal/pkg/A-comunication/client"
	commands "github.com/apfgijon/cartones/internal/pkg/B-commands/commands/commandsv1"
	"github.com/apfgijon/cartones/internal/pkg/C-style/prov"
	"github.com/apfgijon/cartones/internal/pkg/D-filesystem/filesystem"
	"github.com/apfgijon/cartones/pkg/cartongen"
	"github.com/apfgijon/cartones/pkg/covid"
	"github.com/apfgijon/cartones/pkg/pokemon"
	"github.com/google/wire"
)

func InitializeBot(c client.Communication, pokeGame string) (bot.Bot, error) {
	wire.Build(pokemon.NewPokemonImpl,
		covid.NewCovidApi, bot.NewGeneralBot,
		commands.NewCommandImpl,
		prov.NewMessageProoviderv1,
		cartongen.NewCartonv1,
		filesystem.NewFileProvider)
	return &bot.Generalbot{}, nil
}
