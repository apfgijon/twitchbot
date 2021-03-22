package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type PokeapiImpl struct {
	TransMoves map[string]string
	Movesfile  string
}

func (pI *PokeapiImpl) Build() {
	pI.Movesfile = "SpanishMoves.json"
	pI.TransMoves = make(map[string]string)
	pI.initMoves()
}

func (pI *PokeapiImpl) initMoves() {
	transMoves := make(map[string]string)

	phrasesRaw, err := ioutil.ReadFile(pI.Movesfile)
	if err != nil {
		os.Exit(1)
	}

	json.Unmarshal(phrasesRaw, &transMoves)

	pI.TransMoves = transMoves

}

func (pI *PokeapiImpl) PokeMoves(pokeS string, gameVersion string) map[int]string {
	poke, _ := pI.checkIsAndRetPokemon(pokeS)

	return pI.getGameMoves(poke, gameVersion)
}

func (pI *PokeapiImpl) PokeMovesFormatted(pokeS string, gameVersion string) string {
	poke, _ := pI.checkIsAndRetPokemon(pokeS)

	rawMoves := pI.getGameMoves(poke, gameVersion)

	if len(rawMoves) == 0 {
		return ""
	}

	return pI.formatMoves(rawMoves, poke.Name)
}

func (pI *PokeapiImpl) Types(pokeS string) string {
	p, _ := pI.checkIsAndRetPokemon(strings.ToLower(pokeS))
	if p.Name == "" {
		return ""
	}
	typo := p.Types
	typos := ""
	for _, v := range typo {
		TypeName := v.Type.Name
		Tipo, _ := pokeapi.Type(TypeName)
		for _, t := range Tipo.Names {
			if t.Language.Name == "es" {
				TypeName = t.Name
			}
		}

		typos += TypeName + " "
	}
	return typos
}

func (pI *PokeapiImpl) CaptureRate(pokeS string) int {
	p, _ := pokeapi.PokemonSpecies(strings.ToLower(pokeS))
	return p.CaptureRate
}

func (pI *PokeapiImpl) Stats(pokeS string) string {
	p, _ := pI.checkIsAndRetPokemon(strings.ToLower(pokeS))

	stats := p.Name + ": "

	for _, v := range p.Stats {
		stats += v.Stat.Name + "=" + strconv.Itoa(v.BaseStat) + " | "

	}
	return stats
}

func (pI *PokeapiImpl) sortMoves(moves map[int]string) []int {
	levels := make([]int, 0, len(moves))

	for i := range moves {
		levels = append(levels, i)
	}
	sort.Ints(levels)
	return levels
}

func (pI *PokeapiImpl) formatMoves(moves map[int]string, name string) string {

	index := pI.sortMoves(moves)

	formatPokemonMoves := name + ":"
	for _, v := range index {
		if v <= 1 {
			formatPokemonMoves += " lvl:" + "1" + "->" + moves[v] + " 🤙 "
		} else {
			formatPokemonMoves += " lvl:" + strconv.Itoa(v) + "->" + moves[v] + " 🤙 "
		}
	}

	return formatPokemonMoves
}

func (pI *PokeapiImpl) getGameMoves(poke structs.Pokemon, gameVersion string) map[int]string {
	moves := make(map[int]string)

	rawMoves := poke.Moves

	i := 1

	for _, completeMove := range rawMoves {
		for _, versionMove := range completeMove.VersionGroupDetails {
			if versionMove.VersionGroup.Name == gameVersion {
				MoveName := completeMove.Move.Name

				MoveName = pI.getSpanishMove(MoveName)
				if versionMove.LevelLearnedAt != 0 {
					if versionMove.LevelLearnedAt == 1 {
						moves[i] = MoveName
						i--
					} else {
						moves[versionMove.LevelLearnedAt] = MoveName
					}
				}
			}
		}

	}
	return moves
}

type Item struct {
	Name string
	URL  string
}

func (pI *PokeapiImpl) PokeEvos(pokeS string) string {

	pS, _ := pokeapi.PokemonSpecies(strings.ToLower(pokeS))

	allurl := strings.Split(pS.EvolutionChain.URL, "/")
	if len(allurl) < 3 {
		return ""
	}
	evos, _ := pokeapi.EvolutionChain(allurl[len(allurl)-2])
	evosString := evos.Chain.Species.Name

	if len(evos.Chain.EvolvesTo) == 0 {
		return evosString + " no tiene evoluciones :("
	}

	if len(evos.Chain.EvolvesTo[0].EvolutionDetails) != 0 {

		Newevos := evos.Chain.EvolvesTo
		for Newevos[0].Species.Name != "" {
			evosString += " -> "

			for i := 0; i < len(Newevos); i++ {
				if i != 0 {
					evosString += " | "
				}
				if Newevos[i].EvolutionDetails[0].MinLevel != 0 {
					evosString += "lvl: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].MinLevel)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].TradeSpecies != nil {
					evosString += "Tradear: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].TradeSpecies)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].HeldItem != nil {
					itemRaw, _ := json.Marshal(Newevos[i].EvolutionDetails[0].HeldItem)

					var item Item

					json.Unmarshal(itemRaw, &item)
					evosString += "Tener Item puesto: " + item.Name
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].Item != nil {
					itemRaw, _ := json.Marshal(Newevos[i].EvolutionDetails[0].Item)

					var item Item

					json.Unmarshal(itemRaw, &item)

					evosString += "item: " + item.Name
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].KnownMove != nil {
					itemRaw, _ := json.Marshal(Newevos[i].EvolutionDetails[0].KnownMove)
					var item Item

					json.Unmarshal(itemRaw, &item)
					evosString += "Saber movimiento: " + pI.getSpanishMove(item.Name)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].KnownMoveType != nil {
					itemRaw, _ := json.Marshal(Newevos[i].EvolutionDetails[0].KnownMoveType)
					var item Item

					json.Unmarshal(itemRaw, &item)

					var name string

					Tipo, _ := pokeapi.Type(item.Name)
					for _, t := range Tipo.Names {
						if t.Language.Name == "es" {
							name = t.Name
						}
					}
					evosString += "Saber movimiento de tipo: " + name
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].Location != nil {
					itemRaw, _ := json.Marshal(Newevos[i].EvolutionDetails[0].Location)

					var item Item

					json.Unmarshal(itemRaw, &item)
					evosString += "Estar en: " + fmt.Sprint(item.Name)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].MinAffection != nil {
					evosString += "Por Afecto: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].MinAffection)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].MinBeauty != nil {
					evosString += "Por Belleza: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].MinBeauty)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].MinHappiness != nil {
					evosString += "Por Felicidad: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].MinHappiness)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].NeedsOverworldRain != false {
					evosString += "Por lluvia: "
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].PartySpecies != nil {
					evosString += "Por especies en el equipo: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].PartySpecies)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].PartyType != nil {
					evosString += "Por tipos en el equipo: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].PartyType)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].RelativePhysicalStats != nil {
					evosString += "Por estadística: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].RelativePhysicalStats)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].TimeOfDay != "" {
					evosString += "Por horas de la consola: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].TimeOfDay)
					evosString += " " + Newevos[i].Species.Name
				} else if Newevos[i].EvolutionDetails[0].Gender != nil {
					evosString += "Por genero: " + fmt.Sprint(Newevos[i].EvolutionDetails[0].Gender)
					evosString += " " + Newevos[i].Species.Name
				} else {

					evosString += fmt.Sprint(Newevos[i].EvolutionDetails[0].Trigger.Name) + ":"
					evosString += " " + Newevos[i].Species.Name
				}
			}

			if len(Newevos[0].EvolvesTo) != 0 {
				rawEvos, _ := json.Marshal(Newevos[0].EvolvesTo)
				json.Unmarshal(rawEvos, &Newevos)
			} else {
				Newevos[0].Species.Name = ""
			}

		}
	}

	return evosString
}

func (pI *PokeapiImpl) getSpanishMove(MoveName string) string {
	return pI.TransMoves[MoveName]
}

func (pI *PokeapiImpl) checkIsAndRetPokemon(args string) (structs.Pokemon, error) {
	p, err := pokeapi.Pokemon(args)

	return p, err
}
