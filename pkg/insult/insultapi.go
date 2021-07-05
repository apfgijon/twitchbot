package insult

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type InsultProvider struct {
	apiUrl string
}
type insultData struct {
	Number    string `json:"number"`
	Language  string `json:"language"`
	Insult    string `json:"insult"`
	Created   string `json:"created"`
	Shown     string `json:"shown"`
	Createdby string `json:"createdby"`
	Active    string `json:"active"`
	Comment   string `json:"comment"`
}

func (i *InsultProvider) Build() {
	i.apiUrl = "https://evilinsult.com/generate_insult.php?lang=es&type=json"
}
func (i *InsultProvider) Insult() string {
	resp, err := http.Get(i.apiUrl)

	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var ret insultData

	json.Unmarshal(body, &ret)

	return ret.Insult
}
