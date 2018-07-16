package main

import (
	"fmt"
)

func changePointer(ptr *int) {
	*ptr = 20
}

func main() {
	n := 1
	changePointer(&n)
	// Prints 20
	fmt.Println(n)

	// Kinda dangerous will come back to this
}
