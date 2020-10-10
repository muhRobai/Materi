package main

import "log"

func main() {
	var cars = []string{"BMW", "TOYOTA", "HONDA", "YAMAHA"}

	log.Println(cars[0])

	var name = "bambang pamungkas"
	log.Println(name[8:])
	log.Println(name[:7])
	log.Println(name[5:11])

	carsitem := cars[:2]

	log.Println(carsitem)      // BMW TOYOTA
	log.Println(len(carsitem)) // panjang dari carsItem 2
	log.Println(cap(carsitem)) // panjang sebenarnya dari cardsItem yang merupakan potongan dari cars

	carsitem = append(carsitem, "DAIHATSU") // saya menambahkan daihatsu di slice carsItem

	log.Println(carsitem)
	log.Println(len(carsitem))
	log.Println(cap(carsitem))
	log.Println(cars)

	var new = []string{"omama"}
	var newL = []string{"olala", "balala"}

	newAll := copy(newL, new)
	log.Println(new)
	log.Println(newL)
	log.Println(newAll)
}
