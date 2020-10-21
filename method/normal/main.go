package main

import "log"

type student struct {
	name  string
	grade int
}

func (c student) getStudent() student {
	return c
}

func (c student) setSudent(name string, grade int) {
	c.name = name
	c.grade = grade
}

func main() {
	var s1 student

	s1.setSudent("bambang", 30)

	log.Println(s1.getStudent())
}
