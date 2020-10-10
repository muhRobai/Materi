package main

import "fmt"

func main(){
	// komentar
	/*
	fmt.Println("hello world0") */
	
	var item string
	items := ""
	
	var isError bool
	isSuccess := true
	
	name := new(string)
	value := "olala"
	name = &value
	fmt.Println(name)
	fmt.Println(*name)
	fmt.Printf("hello-world-2 %s %s %v %v", item, items, isError, isSuccess)
}  