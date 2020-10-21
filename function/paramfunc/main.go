package main

import (
	"log"
	"strings"
)

// func as a paramater
func filter(data []string) []string {
	var result []string
	for _, p := range data {
		filtered := callback(p)
		if filtered {
			result = append(result, p)
		}
	}

	return result
}

func callback(each string) bool {
	return strings.Contains(each, "o")
}

func main() {
	var data = []string{"omama", "olala", "ethan"}
	var arrData = filter(data)

	log.Println(arrData)
}
