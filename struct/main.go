package main

import (
	"log"
)

// declare struct
type student struct {
	name   string
	kelas  string
	number int
}

func main() {
	// pointer struct
	var c *student

	// add value pointer struct
	c = &student{
		name: "olala",
	}

	d := c
	log.Println(d.name)
	c.name = "sasas"
	log.Println(c.name)
	log.Println(d.name)
}
