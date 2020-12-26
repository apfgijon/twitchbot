package cartongen

import (
	"fmt"
	"math/rand"
)

func GenerateCarton() string {
	var carton [3][9]string

	fillWithRandomNumbers(&carton)

	put0sIntoCarton(&carton)

	ret := stripCarton(carton)

	return ret
}

func fillWithRandomNumbers(carton *[3][9]string) {
	noDuplicateNumbers := []int{}
	var cartonInt [3][9]int

	for row := 0; row < len(carton); row++ {
		for col := 0; col < len(carton[row]); col++ {

			number := generateRandomNumberCarton()
			for contains(noDuplicateNumbers, number) {
				number = generateRandomNumberCarton()
			}
			cartonInt[row][col] = number
			if number < 10 {
				carton[row][col] = "0" + fmt.Sprintf("%d", number)
			} else {
				carton[row][col] = fmt.Sprintf("%d", number)
			}
			noDuplicateNumbers = append(noDuplicateNumbers, number)
		}
	}
}

func put0sIntoCarton(carton *[3][9]string) {
	for row := 0; row < len(carton); row++ {
		XPosition := []int{}
		for i := 0; i < 4; i++ {
			number := rand.Intn(9)
			for contains(XPosition, number) {
				number = rand.Intn(9)
			}
			XPosition = append(XPosition, number)
		}
		for i := 0; i < len(XPosition); i++ {
			carton[row][XPosition[i]] = "XX"
		}
	}
}

func contains(s []int, number int) bool {
	for _, v := range s {
		if v == number {
			return true
		}
	}

	return false
}

func generateRandomNumberCarton() int {
	return rand.Intn(99) + 1
}

func stripCarton(carton [3][9]string) string {
	InitString := " |"
	for row := 0; row < len(carton); row++ {
		for col := 0; col < len(carton[row]); col++ {
			InitString = InitString + carton[row][col] + "|"
		}
		if row != 2 {
			InitString += " |"
		}
	}
	return InitString
}
