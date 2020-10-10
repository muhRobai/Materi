package main

import "log"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	log.Println(numberA)
	log.Println(&numberA)

	log.Println(*numberB)
	log.Println(numberB)

	changePointerValue(numberB)
}

func changePointerValue(item *int) {
	var numberA int = *item
	var numberB *int = &numberA

	log.Println(numberA)
	log.Println(&numberA)

	log.Println(*numberB)
	log.Println(numberB)

	numberA = 10

	log.Println(*numberB)
	log.Println(numberB)
}
