package main

import (
	"fmt"
)

func main() {

	// Use make to make a map
	m := make(map[string]int)

	// Add elements like array
	m["pizza"] = 1
	m["taco"] = 2
	fmt.Println(m)

	// Remove elements with delete
	delete(m, "pizza")
	fmt.Println(m)

	// Can immediately instantiate a map with values
	instantiated := map[string]string{"fish": "friend", "kelp": "food"}
	fmt.Println(instantiated)

	// Second value in a map retreival tells me if that key exists
	_, valueExists := instantiated["shark"]
	fmt.Println(valueExists)
	_, valueExists = instantiated["fish"]
	fmt.Println(valueExists)
}
