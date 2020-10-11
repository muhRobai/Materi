package main

import (
	. "acc-modifier/case"
	api "acc-modifier/library"
	"log"
)

func main() {
	api.SayHello()
	var s1 = api.Items{
		Number: "10",
		Value:  10,
	}

	log.Println("student item", Student.Number)
	log.Println("student item", Student.Value)

	log.Println(s1)

	sayHello("bambang")
}

// how to running go run main.go part.go
// all package main must be run
