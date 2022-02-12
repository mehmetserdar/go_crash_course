package main

import (
	"fmt"
)

func main() {
	ids := []int{34, 65, 26, 57, 84, 7}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("Index: %d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf(" ID: %d\n", id)
	}

	//Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with Map
	emails := map[string]string{"bob": "bob@gmail.com", "john": "john@gmail.com", "mike": "mike@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
