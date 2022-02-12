package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + "and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// init Person using struct
	person1 := Person{firstName: "Jane", lastName: "Smith", city: "London", gender: "f", age: 25}
	// alternative
	person2 := Person{"Bob", "Jackson", "New York", "m", 30}

	// fmt.Println(person1)
	// fmt.Println(person2)

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Peters")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
