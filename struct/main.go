package main

import (
	"fmt"
	"log"
)

type student struct {
	name   string
	kelas  string
	number int
}

func main() {

	var item student
	item.name = "bambang"
	item.kelas = "1"

	fmt.Println(item.name)
	fmt.Println(item.kelas)
	fmt.Println(item.number)

	data := student{
		name:   "saskia",
		kelas:  "31",
		number: 10,
	}

	log.Println(data)

	var c *student

	c = &student{
		name: "olala",
	}

	log.Println(c)
}
