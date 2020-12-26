package covid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type FullCovidDataArray struct {
	Data []CovidData `json:"data"`
}

type CovidData struct {
	Date          string  `json:"date"`
	Confirmed     int     `json:"confirmed"`
	Deaths        int     `json:"deaths"`
	Recovered     int     `json:"recovered"`
	ConfirmedDiff int     `json:"confirmed_diff"`
	DeathsDiff    int     `json:"deaths_diff"`
	RecoveredDiff int     `json:"recovered_diff"`
	LastUpdate    string  `json:"last_update"`
	Active        int     `json:"active"`
	ActiveDiff    int     `json:"active_diff"`
	FatalityRate  float64 `json:"fatality_rate"`
	Region        struct {
		Iso      string        `json:"iso"`
		Name     string        `json:"name"`
		Province string        `json:"province"`
		Lat      string        `json:"lat"`
		Long     string        `json:"long"`
		Cities   []interface{} `json:"cities"`
	} `json:"region"`
}

func getRawData() []byte {
	resp, err := http.Get("https://covid-api.com/api/reports?iso=ESP")

	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return body
}

func GetCovidCases() *FullCovidDataArray {
	var retData *FullCovidDataArray
	json.Unmarshal(getRawData(), &retData)
	return retData
}

func GetCovidCasesForProvince(province string) CovidData {
	var provinceCovidData CovidData
	province = strings.ToLower(province)
	for _, data := range GetCovidCases().Data {
		if strings.ToLower(data.Region.Province) == province {
			provinceCovidData = data
		}
	}
	return provinceCovidData
}

func GetCovidCasesSpain() (int, int) { //(casos,muertos)
	casos := 0
	muertos := 0
	for _, data := range GetCovidCases().Data {
		casos += data.ConfirmedDiff
		muertos += data.DeathsDiff
	}
	return casos, muertos
}
