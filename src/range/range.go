package main

import (
	"fmt"
)

func main() {
	// iterate over array
	names := [5]string{"bob", "sally", "fernando", "billy", "vanessa"}
	for i, name := range names {
		fmt.Print(i, ":", name, " ")
	}
	fmt.Println()

	// iterate over map
	carMap := map[string]string{"ford": "gt", "ferrari": "458", "lamborghini": "hurrican"}
	for k, v := range carMap {
		fmt.Println("The all new", k, v)
	}
}
