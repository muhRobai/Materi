package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	fmt.Print("Masukkan nilai input: ")
	fmt.Scan(&input)

	mid := float64(input / 2)
	if input%2 == 0 {
		fmt.Println("Nilai input tidak boleh genap")
	} else {
		for i := 0; i < input; i++ {
			pattern := ""

			for j := 0; j < input; j++ {
				if j == 0 || j == input-1 {
					pattern += "* "
				} else if float64(i) == math.Floor(mid) {
					pattern += "* "
				} else {
					pattern += "= "
				}
			}
			fmt.Println(pattern)
		}
	}
}
