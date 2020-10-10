package main

import (
	"log"
)

const pi = 3.14

func main() {
	printMessage("hello wolrd")
	log.Println(getRange(10, 11))
	log.Println(calculate(1.5))

	log.Println(variadicCalculate(1, 1, 4, 2, 5, 7, 8, 9, 0, 12))
}

func printMessage(message string) {
	log.Println(message)
}

func getRange(min, max int) int {
	return max * min
}

func calculate(r float64) (float64, float64) {
	var area = pi * (r * r)

	var circle = 2 * pi * r

	return area, circle
}

func variadicCalculate(number ...int) float64 {
	var total int = 0
	for _, number := range number {
		total += number
	}

	var avg = float64(total) / float64(len(number))

	return avg
}
