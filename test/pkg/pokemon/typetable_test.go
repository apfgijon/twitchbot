package test

import (
	"testing"

	"github.com/apfgijon/cartones/pkg/pokemon"
)

func TestPokemonTable(t *testing.T) {
	expectedBulbasaurTableType := make(map[string]int)
	expectedBulbasaurTableType["Agua"] = -1
	expectedBulbasaurTableType["Electrico"] = -1
	expectedBulbasaurTableType["Fuego"] = 1
	expectedBulbasaurTableType["Hada"] = -1
	expectedBulbasaurTableType["Hielo"] = 1
	expectedBulbasaurTableType["Lucha"] = -1
	expectedBulbasaurTableType["Planta"] = -2
	expectedBulbasaurTableType["Psiquico"] = 1
	expectedBulbasaurTableType["Volador"] = 1

	pokemonProvider, _ := pokemon.NewPokemonImpl("HG")

	_, obtainedBulbasaurTableType, _ := pokemonProvider.TypeTablePokemon("bulbasaur")

	for i, v := range expectedBulbasaurTableType {
		if v != obtainedBulbasaurTableType[i] {
			t.Errorf("Type %s: Expected: %d, Returned: %d", i, v, obtainedBulbasaurTableType[i])
		}
	}
}

func Test2TypeTable(t *testing.T) {
	expectedVenenoPlantaTableType := make(map[string]int)
	expectedVenenoPlantaTableType["Agua"] = -1
	expectedVenenoPlantaTableType["Electrico"] = -1
	expectedVenenoPlantaTableType["Fuego"] = 1
	expectedVenenoPlantaTableType["Hada"] = -1
	expectedVenenoPlantaTableType["Hielo"] = 1
	expectedVenenoPlantaTableType["Lucha"] = -1
	expectedVenenoPlantaTableType["Planta"] = -2
	expectedVenenoPlantaTableType["Psiquico"] = 1
	expectedVenenoPlantaTableType["Volador"] = 1

	pokemonProvider, _ := pokemon.NewPokemonImpl("HG")

	_, obtainedBulbasaurTableType, _ := pokemonProvider.TypeTablePokemon("Planta Veneno")

	for i, v := range expectedVenenoPlantaTableType {
		if v != obtainedBulbasaurTableType[i] {
			t.Errorf("Type %s: Expected: %d, Returned: %d", i, v, obtainedBulbasaurTableType[i])
		}
	}
}

func Test1TypeTable(t *testing.T) {
	expectedVenenoPlantaTableType := make(map[string]int)
	expectedVenenoPlantaTableType["Agua"] = -1
	expectedVenenoPlantaTableType["Bicho"] = 1
	expectedVenenoPlantaTableType["Electrico"] = -1
	expectedVenenoPlantaTableType["Fuego"] = 1
	expectedVenenoPlantaTableType["Hielo"] = 1
	expectedVenenoPlantaTableType["Planta"] = -1
	expectedVenenoPlantaTableType["Tierra"] = -1
	expectedVenenoPlantaTableType["Veneno"] = 1
	expectedVenenoPlantaTableType["Volador"] = 1

	pokemonProvider, _ := pokemon.NewPokemonImpl("HG")

	_, obtainedBulbasaurTableType, _ := pokemonProvider.TypeTablePokemon("Planta")

	for i, v := range expectedVenenoPlantaTableType {
		if v != obtainedBulbasaurTableType[i] {
			t.Errorf("Type %s: Expected: %d, Returned: %d", i, v, obtainedBulbasaurTableType[i])
		}
	}
}
