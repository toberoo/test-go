package main

import (
	"fmt"
)

// squares the provided number
func square(a int) int {
	return a * a
}

// multiple returns
func isEvenIsOdd(n int) (bool, bool) {
	return n%2 == 0, n%2 == 1
}

// n param functions
func spacedPrint(items ...string) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

// Closures and higher order function
func addByX(i int) func(int) int {
	return func(b int) int {
		return i + b
	}
}

func main() {
	const four = 4
	fmt.Println(square(four))

	// Handle multiple returns
	_, isOdd := isEvenIsOdd(four)
	fmt.Println(isOdd)

	// Multiple inputs
	spacedPrint("taco", "fish", "pizza")

	// Can also spread the array to pass in
	spacedPrint([]string{"carry", "me", "home"}...)

	// Create the add by four func
	addByFourFunc := addByX(four)
	fmt.Println(addByFourFunc(2))
}
