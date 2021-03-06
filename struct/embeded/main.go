package main

import "log"

// initial struct
type student struct {
	name   string
	kelas  string
	number int
}

// embeded struct call struct on struct
type class struct {
	list  []*student
	total string
	// avg   student
}

func main() {
	var items []*student
	item := student{
		name:  "omama",
		kelas: "olala",
	}

	items = append(items, &item)

	log.Println(items[0])

	// anonymoust struct and using tag on struct
	var p = struct {
		Jumlah string `json:"jumlah"`
		Total  string `xml:"total"`
	}{"12", "30"}

	log.Println(p)
}
