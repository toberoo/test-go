package main

import "fmt"

func main() {

	// Slices are made with make
	s := make([]string, 2)
	fmt.Println("empty slice:", s)

	// Can acess and set like array
	s[0] = "hello"
	s[1] = "world"
	fmt.Println(s[0], s[1])

	// I don't like how we have to use len to get the length
	fmt.Println("this slice is", len(s), "long")

	// We can also append values immutably
	appended := append(s, "taco")
	fmt.Println("new array is", appended)
	fmt.Println("old array is still", s)

	// We can copy slices. I don't like how
	// this function doesn't have a return value.
	copiedArray := make([]string, 2)
	copy(copiedArray, s)

	// Instatiate slice with values
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(a, "with all values")

	// We can also do that python stuff where we only tak a portion of the slices

	// skip first element
	partial := a[1:]
	fmt.Println(partial, "beginning slice")

	// Skip last two elements
	partial = a[:3]
	fmt.Println(partial, "ending slice")

	// Middle slice
	partial = a[2:3]
	fmt.Println(partial, "middle slice")

	// Slice of Slices
	goodOle := make([][]int, 3)

	for i := 0; i < len(goodOle); i++ {
		for j := 0; j < len(goodOle[i]); j++ {
			fmt.Println(i)
		}
	}

	fmt.Println(goodOle, "double slice")
}
