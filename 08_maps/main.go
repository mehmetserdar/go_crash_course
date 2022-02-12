package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[int]string)

	// // Asign kv
	// emails[1] = "bob@gmail.com"
	// emails[2] = "john@gmail.com"
	// emails[3] = "mike@gmail.com"

	// Declare map and kv

	emails := map[int]string{1: "bob@gmail.com", 2: "john@gmail.com", 3: "mike@gmail.com"}
	emails[4] = "jane@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails[1])

	// Delete from Map
	delete(emails, 1)
	fmt.Println(emails)
}
