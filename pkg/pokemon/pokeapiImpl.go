package pokemon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

type PokeapiImpl struct {
	TransMoves map[string]string
	TransTypes map[string]string
	Movesfile  string
	Typesfile  string
	game       string
}

func NewPokemonImpl(game string) (PokeInfo, error) {
	var gameSelected string
	switch game {
	case "HG":
		gameSelected = "heartgold-soulsilver"
		break
	default:
		gameSelected = "heartgold-soulsilver"
	}

	pI := &PokeapiImpl{
		game: gameSelected,
	}
	pI.Movesfile = "SpanishMoves.json"
	pI.TransMoves = make(map[string]string)
	pI.Typesfile = "SpanishTypes.json"
	pI.TransTypes = make(map[string]string)
	pI.initMoves()

	return pI, nil
}

func (pI *PokeapiImpl) initMoves() {
	transMoves := make(map[string]string)

	movesRaw, err := ioutil.ReadFile(pI.Movesfile)
	if err != nil {
		os.Exit(1)
	}

	json.Unmarshal(movesRaw, &transMoves)

	pI.TransMoves = transMoves

	transTypes := make(map[string]string)

	typesRaw, err := ioutil.ReadFile(pI.Typesfile)
	if err != nil {
		os.Exit(1)
	}

	json.Unmarshal(typesRaw, &transTypes)

	pI.TransTypes = transTypes

}

func (pI *PokeapiImpl) PokeMoves(pokeS string) map[int]string {
	poke, _ := pI.checkIsAndRetPokemon(pokeS)

	return pI.getGameMoves(poke)
}

func (pI *PokeapiImpl) PokeMovesFormatted(pokeS string) string {
	poke, err := pI.checkIsAndRetPokemon(pokeS)

	if err != nil {
		return ""
	}

	rawMoves := pI.getGameMoves(poke)

	if len(rawMoves) == 0 {
		return ""
	}

	return pI.formatMoves(rawMoves, poke.Name)
}

func (pI *PokeapiImpl) Types(pokeS string) string {
	p, err := pI.checkIsAndRetPokemon(strings.ToLower(pokeS))

	if err != nil {
		return ""
	}

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
	p, err := pI.checkIsAndRetPokemon(strings.ToLower(pokeS))

	if err != nil {
		return ""
	}
	stats := p.Name + ": "

	for _, v := range p.Stats {
		stats += v.Stat.Name + "=" + strconv.Itoa(v.BaseStat) + " | "

	}
	return stats
}

func (pI *PokeapiImpl) Peso(pokeS string) string {
	p, err := pI.checkIsAndRetPokemon(strings.ToLower(pokeS))

	if err != nil {
		return ""
	}
	stats := p.Name + ": "

	stats += "peso: " + strconv.FormatFloat(float64(p.Weight)*0.1, 'f', 1, 64) + " kg "
	stats += "altura: " + strconv.Itoa(p.Height) + " m"

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

func (pI *PokeapiImpl) PokeRandom() string {

	l, _ := pokeapi.Resource("pokemon", 0, 386)

	return l.Results[rand.Intn(len(l.Results))].Name
}

func (pI *PokeapiImpl) getGameMoves(poke structs.Pokemon) map[int]string {
	moves := make(map[int]string)

	rawMoves := poke.Moves

	i := 1

	for _, completeMove := range rawMoves {
		for _, versionMove := range completeMove.VersionGroupDetails {
			if versionMove.VersionGroup.Name == pI.game {
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

func (pI *PokeapiImpl) getGameMovesMT(poke structs.Pokemon) []string {
	moves := make([]string, 0)

	rawMoves := poke.Moves

	for _, completeMove := range rawMoves {
		for _, versionMove := range completeMove.VersionGroupDetails {
			if versionMove.VersionGroup.Name == pI.game {
				MoveName := completeMove.Move.Name

				MoveName = pI.getSpanishMove(MoveName)
				if versionMove.MoveLearnMethod.Name == "machine" {

					moves = append(moves, MoveName)
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

func (pI *PokeapiImpl) TypeTable(typo string) string {

	TypeS := pI.getEnglishType(typo)

	if TypeS == "" {
		return TypeS
	}

	Type, _ := pokeapi.Type(TypeS)

	ret := "Tipo: " + pI.getSpanishType(Type.Name) + " | "

	ret += pI.formatTypesTo(Type)

	ret += " | "

	ret += pI.formatTypesFrom(Type)

	return ret
}

func (pI *PokeapiImpl) TypeTablePokemon(poke string) (int, map[string]int, map[string]int) {

	esPokemon := 0

	if poke == "" {
		return esPokemon, nil, nil
	}
	pokemonT := pI.Types(poke)
	pokemonT = strings.ReplaceAll(pokemonT, "á", "a")
	pokemonT = strings.ReplaceAll(pokemonT, "ó", "o")
	pokemonT = strings.ReplaceAll(pokemonT, "í", "i")
	pokemonT = strings.ReplaceAll(pokemonT, "ú", "u")
	pokemonT = strings.ReplaceAll(pokemonT, "é", "e")

	var types []string
	if pokemonT != "" {
		types = strings.Split(pokemonT, " ")
		esPokemon = 1
	} else {
		types = strings.Split(poke, " ")
		if len(types) == 1 {
			esPokemon = 2
		} else {
			esPokemon = 3
		}
	}

	tablatiposFrom := make(map[string]int)
	tablatiposTo := make(map[string]int)

	for _, t := range types {
		TypeS := pI.getEnglishType(t)
		Type, _ := pokeapi.Type(TypeS)

		if len(Type.DamageRelations.DoubleDamageFrom) > 0 {
			for _, v := range Type.DamageRelations.DoubleDamageFrom {
				tablatiposFrom[pI.getSpanishType(v.Name)]++
			}
		}

		if len(Type.DamageRelations.HalfDamageFrom) > 0 {
			for _, v := range Type.DamageRelations.HalfDamageFrom {
				itemRaw, _ := json.Marshal(v)

				var item Item

				json.Unmarshal(itemRaw, &item)
				tablatiposFrom[pI.getSpanishType(item.Name)]--
			}
		}

		if len(Type.DamageRelations.NoDamageFrom) > 0 {
			for _, v := range Type.DamageRelations.NoDamageFrom {
				tablatiposFrom[pI.getSpanishType(v.Name)] = -100
			}
		}

		if len(Type.DamageRelations.DoubleDamageTo) > 0 {
			for _, v := range Type.DamageRelations.DoubleDamageTo {
				itemRaw, _ := json.Marshal(v)

				var item Item

				json.Unmarshal(itemRaw, &item)
				tablatiposTo[pI.getSpanishType(item.Name)]++
			}
		}

		if len(Type.DamageRelations.HalfDamageTo) > 0 {
			for _, v := range Type.DamageRelations.HalfDamageTo {
				tablatiposTo[pI.getSpanishType(v.Name)]--
			}
		}

		if len(Type.DamageRelations.NoDamageTo) > 0 {
			for _, v := range Type.DamageRelations.NoDamageTo {
				tablatiposTo[pI.getSpanishType(v.Name)] = -100
			}
		}
	}

	return esPokemon, tablatiposFrom, tablatiposTo
}

func (pI *PokeapiImpl) TypeTableFrom(typo string) string {

	TypeS := pI.getEnglishType(typo)

	if TypeS == "" {
		return TypeS
	}

	Type, _ := pokeapi.Type(TypeS)

	ret := "Tipo: " + pI.getSpanishType(Type.Name) + " | "

	ret += pI.formatTypesFrom(Type)

	return ret
}

func (pI *PokeapiImpl) TypeTableTo(typo string) string {

	TypeS := pI.getEnglishType(typo)

	if TypeS == "" {
		return TypeS
	}

	Type, _ := pokeapi.Type(TypeS)

	ret := "Tipo: " + pI.getSpanishType(Type.Name) + " | "

	ret += pI.formatTypesTo(Type)

	return ret
}

func (pI *PokeapiImpl) formatTypesFrom(Type structs.Type) string {
	format := ""

	if len(Type.DamageRelations.DoubleDamageFrom) > 0 {
		format += " LU JODE: "
		for _, v := range Type.DamageRelations.DoubleDamageFrom {
			format += pI.getSpanishType(v.Name) + ", "
		}
		format = format[:len(format)-2]
	}

	if len(Type.DamageRelations.HalfDamageFrom) > 0 {
		format += " | "
		format += " RESISTE: "

		for _, v := range Type.DamageRelations.HalfDamageFrom {
			itemRaw, _ := json.Marshal(v)

			var item Item

			json.Unmarshal(itemRaw, &item)
			format += pI.getSpanishType(item.Name) + ", "
		}
		format = format[:len(format)-2]
	}

	if len(Type.DamageRelations.NoDamageFrom) > 0 {
		format += " | "
		format += " NO LE AFECTA: "
		for _, v := range Type.DamageRelations.NoDamageFrom {
			format += pI.getSpanishType(v.Name) + ", "
		}
		format = format[:len(format)-2]
	}

	return format
}

func (pI *PokeapiImpl) formatTypesTo(Type structs.Type) string {
	format := ""

	if len(Type.DamageRelations.DoubleDamageTo) > 0 {
		format += " EFECTIVO CONTRA: "
		for _, v := range Type.DamageRelations.DoubleDamageTo {
			itemRaw, _ := json.Marshal(v)

			var item Item

			json.Unmarshal(itemRaw, &item)
			format += pI.getSpanishType(item.Name) + ", "
		}
		format = format[:len(format)-2]
	}

	if len(Type.DamageRelations.HalfDamageTo) > 0 {
		format += " | "
		format += " PUTA MIERDA CONTRA: "
		for _, v := range Type.DamageRelations.HalfDamageTo {
			format += pI.getSpanishType(v.Name) + ", "
		}
		format = format[:len(format)-2]
	}

	if len(Type.DamageRelations.NoDamageTo) > 0 {
		format += " | "
		format += " NO EFECTIVO CONTRA: "
		for _, v := range Type.DamageRelations.NoDamageTo {
			format += pI.getSpanishType(v.Name) + ", "
		}
		format = format[:len(format)-2]
	}

	return format
}

func (pI *PokeapiImpl) getSpanishType(TypeName string) string {
	return pI.TransTypes[TypeName]
}

func (pI *PokeapiImpl) getEnglishType(TypeName string) string {
	ret := ""
	for i, v := range pI.TransTypes {
		if strings.ToLower(v) == strings.ToLower(TypeName) {
			ret = i
		}
	}
	return ret
}

func (pI *PokeapiImpl) getEnglishAttack(AtackName string) string {
	ret := ""
	for i, v := range pI.TransMoves {
		if strings.ToLower(v) == strings.ToLower(AtackName) {
			ret = i
		}
	}
	return ret
}

func (pI *PokeapiImpl) checkIsAndRetPokemon(args string) (structs.Pokemon, error) {
	p, err := pokeapi.Pokemon(args)

	return p, err
}

func (pI *PokeapiImpl) PP(ataque string) int {
	a := pI.getEnglishAttack(ataque)

	m, _ := pokeapi.Move(a)

	return m.Pp

}
