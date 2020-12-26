package randomsay

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

const file = "frases.json"

func GetRandomPhrase() string {
	phrasesRaw, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	var allPhrases []string

	json.Unmarshal(phrasesRaw, &allPhrases)

	randomNumber := rand.Intn(len(allPhrases))

	return allPhrases[randomNumber]
}

func SetPhrase(frase string) {
	phrasesRaw, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	var allPhrases []string

	json.Unmarshal(phrasesRaw, &allPhrases)

	allPhrases = append(allPhrases, frase)

	phrasesRaw, err = json.MarshalIndent(allPhrases, "", "	")
	if err != nil {
		return
	}

	ioutil.WriteFile(file, phrasesRaw, 0644)
}
