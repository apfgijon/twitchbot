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

type CovidApiImpl struct {
	apiUrl string
}

func NewCovidApi() (CovidInfo, error) {
	cI := &CovidApiImpl{}

	cI.apiUrl = "https://covid-api.com/api/reports?iso=ESP"

	return cI, nil
}

func (cI *CovidApiImpl) getRawData() []byte {
	resp, err := http.Get(cI.apiUrl)

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

func (cI *CovidApiImpl) getCovidCases() *FullCovidDataArray {
	var retData *FullCovidDataArray
	json.Unmarshal(cI.getRawData(), &retData)
	return retData
}

func (cI *CovidApiImpl) GetCovidCasesForProvince(province string) (int, int) {
	var provinceCovidData CovidData
	province = strings.ToLower(province)
	for _, data := range cI.getCovidCases().Data {
		if strings.ToLower(data.Region.Province) == province {
			provinceCovidData = data
		}
	}
	return provinceCovidData.ConfirmedDiff, provinceCovidData.DeathsDiff
}

func (cI *CovidApiImpl) GetCovidCasesSpain() (int, int) { //(casos,muertos)
	casos := 0
	muertos := 0
	for _, data := range cI.getCovidCases().Data {
		casos += data.ConfirmedDiff
		muertos += data.DeathsDiff
	}
	return casos, muertos
}
