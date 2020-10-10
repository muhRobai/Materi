package main

import (
	"fmt"
)

// func StrFunc(req string) (string, error) {
// 	return req, nil
// }

func main() {
	// resp, _ := StrFunc("ommama")
	// log.Println(resp)

	// normal for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for like while
	var l = 0
	for l < 10 {
		fmt.Println(l)
		l++
	}

	// for infinity loop
	var p = 0
	for {
		fmt.Println("angka: ", p)

		if p == 5 {
			break
		}
		p++
	}

	// for like foreach
	var items = []string{"olala", "omama", "balal", "lala"}
	for index, k := range items {
		fmt.Println(index)
		fmt.Println(k)
	}
}
