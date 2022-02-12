package main

import "fmt"

func main() {
	// using var
	//var name = "Mehmet"
	var age int32 = 35
	const isCool = true
	var size float32 = 2.4

	// Shorthand
	//name := "Mehmet"
	//email := "ffgdfg@grr.com"

	name, email := "Mehmet", "meh@yandex.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
