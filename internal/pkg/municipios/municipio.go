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

func get_municipios() ([][]string, error) {

	readerData, err := os.Open("Municipio.csv")
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

	r, _ := ioutil.ReadFile("Municipio.csv")

	rS := string(r[:])

	//rS = strings.ReplaceAll(rS, ",", ".")

	rS = strings.ReplaceAll(rS, ",", ";")

	ioutil.WriteFile("Municipio.csv", []byte(rS), 0)

}

func Valid() {
	r, _ := ioutil.ReadFile("Municipio.csv")
	rS := string(r[:])
	fmt.Println(utf8.ValidString(rS))
}

func HablameSobre(mun string) string {

	ret := ""
	mun = strings.ToLower(mun)
	allMuni, _ := get_municipios()
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
	allMuni, _ := get_municipios()
	for _, data := range allMuni {

		municipioInData := strings.ToLower(data[0])
		if municipioInData == mun {
			ret = data[11]
			break
		}
	}
	return ret
}
