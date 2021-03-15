package municipios

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const municipiosFile string = "Municipio.csv"

func getMunicipios() ([][]string, error) {

	readerData, err := os.Open(municipiosFile)
	if err != nil {
		return nil, err
	}

	readerformatedData := csv.NewReader(readerData)

	readerformatedData.Comma = ';'

	formatedData, err := readerformatedData.ReadAll()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	readerData.Close()
	return formatedData, nil
}

func ConvertToCsv() {

	r, _ := ioutil.ReadFile(municipiosFile)

	rS := string(r[:])

	//rS = strings.ReplaceAll(rS, ",", ".")

	rS = strings.ReplaceAll(rS, ",", ";")

	ioutil.WriteFile(municipiosFile, []byte(rS), 0)

}

func Valid() {
	r, _ := ioutil.ReadFile(municipiosFile)
	rS := string(r[:])
	fmt.Println(utf8.ValidString(rS))
}

func HablameSobre(mun string) string {

	ret := ""
	mun = strings.ToLower(mun)
	allMuni, _ := getMunicipios()
	for _, data := range allMuni {

		municipioInData := strings.ToLower(data[0])
		if municipioInData == mun {
			ret = data[7]
			break
		}
	}
	return ret
}
func QueVer(mun string) string {

	ret := ""
	mun = strings.ToLower(mun)
	allMuni, _ := getMunicipios()
	for _, data := range allMuni {

		municipioInData := strings.ToLower(data[0])
		if municipioInData == mun {
			ret = data[11]
			break
		}
	}
	return ret
}
