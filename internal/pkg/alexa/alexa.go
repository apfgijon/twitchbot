package alexa

import (
	"io/ioutil"
	"strconv"
)

const alexafile = "alexa"

func GetNumber() int {
	numberRaw, err := ioutil.ReadFile(alexafile)
	if err != nil {
		return -1
	}
	number, err := strconv.Atoi(string(numberRaw))

	if err != nil {
		return -1
	}
	number++

	ioutil.WriteFile(alexafile, []byte(strconv.Itoa(number)), 0644)

	return number
}
