package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	// Assign Values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and Assign (shorthand)
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grapes"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
