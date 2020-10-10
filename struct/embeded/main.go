package main

import "log"

type student struct {
	name   string
	kelas  string
	number int
}

type class struct {
	list  []*student
	total string
}

func main() {
	var items []*student
	item := student{
		name:  "omama",
		kelas: "olala",
	}

	items = append(items, &item)

	log.Println(items[0])

	var p = struct {
		Jumlah string `josn:"jumlah"`
		Total  string `json:"total"`
	}{"12", "30"}

	log.Println(p)
}
