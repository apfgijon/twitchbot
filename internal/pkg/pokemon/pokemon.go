package pokemon

import (
	"sort"
	"strconv"

	"github.com/mtslzr/pokeapi-go/structs"
)

func PokeMoves(poke structs.Pokemon, gameVersion string) string {
	rawMoves := getEmeraldMoves(poke, gameVersion)

	if len(rawMoves) == 0 {
		return ""
	}

	return formatMoves(rawMoves, poke.Name)
}

func sortMoves(moves map[int]string) []int {
	levels := make([]int, 0, len(moves))

	for i := range moves {
		levels = append(levels, i)
	}
	sort.Ints(levels)
	return levels
}

func formatMoves(moves map[int]string, name string) string {

	index := sortMoves(moves)

	formatPokemonMoves := name + ":"
	for _, v := range index {
		formatPokemonMoves += " lvl:" + strconv.Itoa(v) + "->" + moves[v] + " 🤙 "
	}

	return formatPokemonMoves
}

func getEmeraldMoves(poke structs.Pokemon, gameVersion string) map[int]string {
	moves := make(map[int]string)

	rawMoves := poke.Moves

	for _, completeMove := range rawMoves {
		for _, versionMove := range completeMove.VersionGroupDetails {
			if versionMove.VersionGroup.Name == gameVersion {
				MoveName := completeMove.Move.Name

				// Move, _ := pokeapi.Move(MoveName)

				// for _, names := range Move.Names {
				// 	if names.Language.Name == "es" {
				// 		MoveName = names.Name
				// 	}
				// }

				moves[versionMove.LevelLearnedAt] = MoveName
			}
		}
	}
	return moves
}
