package main

import (
	"fmt"
)

type cat struct {
	name  string
	color string
}

// Add methods to struct with
// functions that take a refrence to struct
// as a pointer
func (c *cat) meow() {
	fmt.Println("meow")
}

func main() {

	// Creating a struct
	maddie := cat{"maddie", "grey"}
	fmt.Println(maddie.color)

	// Can point to structs like C
	// And change values
	ptr := &maddie
	ptr.color = "red"
	fmt.Println(maddie.color)

	maddie.meow()
}
