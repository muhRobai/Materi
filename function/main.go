package main

import (
	"log"
)

const pi = 3.14

// reminder
// need example for using ponter on function

func main() {
	var arrInt = []int{1, 21, 3, 4, 5, 33, 535, 353, 5, 353, 5, 35, 5353, 535, 35, 353, 535, 35, 35}
	resp, arrItem := findMax(arrInt, 2000)
	log.Println(resp)
	log.Println(arrItem(100))
}

// func as a return
func findMax(number []int, max int) (int, func(int) []int) {
	var res []int
	for _, p := range number {
		if p <= max {
			res = append(res, p)
		}
	}

	return len(res), func(max int) []int {
		var arrInt []int
		for _, q := range res {
			if q > max {
				arrInt = append(arrInt, q)
			}
		}

		return arrInt
	}
}
