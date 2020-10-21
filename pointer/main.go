package main

import "log"

func main() {
	var numberA int = 4
	// int pointer
	var numberB *int = &numberA

	// log.Println(numberA)
	// log.Println(&numberA)

	// // get pointer value
	// log.Println(*numberB)
	// log.Println(numberB)

	changePointerValue(numberB)
}

func changePointerValue(item *int) {
	var numberA int = *item
	var numberB *int = &numberA

	log.Println(numberA, "ini A")
	log.Println(&numberA)

	log.Println(*numberB, "ini B")
	log.Println(numberB)

	log.Println("saya ubah A")
	numberA = 10

	log.Println(*numberB, "nilai B ikut berubah")
	log.Println(numberB)
}
