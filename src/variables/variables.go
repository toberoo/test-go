package main

import "fmt"

func main() {
	var a = 1
	var b = a + 1
	var c = b + 1

	// Prints 6
	fmt.Println(a + b + c)

	// Prints 1 2 3
	fmt.Println(a, b, c)

	// If we don't give it an initial value
	// it requires nice
	var x int
	// Prints 0
	fmt.Println(x)

	x = 2
	// Print 2
	fmt.Println(x)

	// We don't need var if we use :=
	y := "hello"
	fmt.Println(y)

	// Constants can not change refrence.
	const pi = 3.14

}
