package cartongen

import (
	"fmt"
	"math/rand"
)

type Cartonv1 struct {
}

func NewCartonv1() Carton {
	return &Cartonv1{}
}

func (this *Cartonv1) GenerateCarton() string {
	var carton [3][9]string

	this.fillWithRandomNumbers(&carton)

	this.put0sIntoCarton(&carton)

	ret := this.stripCarton(carton)

	return ret
}

func (this *Cartonv1) fillWithRandomNumbers(carton *[3][9]string) {
	noDuplicateNumbers := []int{}
	var cartonInt [3][9]int

	for row := 0; row < len(carton); row++ {
		for col := 0; col < len(carton[row]); col++ {

			number := this.generateRandomNumberCarton()
			for this.contains(noDuplicateNumbers, number) {
				number = this.generateRandomNumberCarton()
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

func (this *Cartonv1) put0sIntoCarton(carton *[3][9]string) {
	for row := 0; row < len(carton); row++ {
		XPosition := []int{}
		for i := 0; i < 4; i++ {
			number := rand.Intn(9)
			for this.contains(XPosition, number) {
				number = rand.Intn(9)
			}
			XPosition = append(XPosition, number)
		}
		for i := 0; i < len(XPosition); i++ {
			carton[row][XPosition[i]] = "XX"
		}
	}
}

func (this *Cartonv1) contains(s []int, number int) bool {
	for _, v := range s {
		if v == number {
			return true
		}
	}

	return false
}

func (this *Cartonv1) generateRandomNumberCarton() int {
	return rand.Intn(99) + 1
}

func (this *Cartonv1) stripCarton(carton [3][9]string) string {
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
