package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var lamp int
	var push int
	fmt.Print("> ")
	fmt.Scanf("%d %d", &lamp, &push)
	item := fmt.Sprintf("%T %v", lamp, lamp)
	strItem := strconv.Itoa(lamp)
	item2 := fmt.Sprintf("%T %v", strItem, strItem)
	log.Println(item)
	log.Println(item2)
	var turnLamp []int

	for i := 1; i <= push; i++ {
		for l := 1; l <= lamp; l++ {
			if i == 1 {
				turnLamp = append(turnLamp, 1)
			}

			if l%i == 0 && i != 1 {
				if int(turnLamp[l-1]) > 0 {
					turnLamp[l-1] = 0
				} else {
					turnLamp[l-1] = 1
				}
			}
		}
	}

	result := 0
	for i := 0; i < len(turnLamp); i++ {
		if int(turnLamp[i]) > 0 {
			result++
		}
	}
	fmt.Print("> ")
	fmt.Println(result)
}
