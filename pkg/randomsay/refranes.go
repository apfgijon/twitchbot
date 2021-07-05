package randomsay

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

const refranesfile = "refranes.json"

func GetRandomRefran() string {
	phrasesRaw, err := ioutil.ReadFile(refranesfile)
	if err != nil {
		return ""
	}
	var allPhrases []string

	json.Unmarshal(phrasesRaw, &allPhrases)

	randomNumber := rand.Intn(len(allPhrases))

	return allPhrases[randomNumber]
}
