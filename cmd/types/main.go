package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mtslzr/pokeapi-go"
)

const file string = "SpanishTypes.json"

type Moves struct {
	MoveEn string
	MoveEs string
}

func main() {
	movesTrans := make(map[string]string)
	for i := 1; i < 19; i++ {

		Move, _ := pokeapi.Type(fmt.Sprint(i))

		var MoveEs string
		var MoveEn string

		for _, names := range Move.Names {
			if names.Language.Name == "es" {
				MoveEs = names.Name
			}
			if names.Language.Name == "en" {
				MoveEn = names.Name
				MoveEn = strings.ToLower(MoveEn)
				MoveEn = strings.Replace(MoveEn, " ", "-", 10)
			}
		}

		movesTrans[MoveEn] = MoveEs
		fmt.Println("En: " + MoveEn + "Es: " + MoveEs)
	}
	movesRaw, err := json.MarshalIndent(movesTrans, "", "	")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ioutil.WriteFile(file, movesRaw, 0644)
}
