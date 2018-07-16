package main

import (
	"fmt"
)

func main() {

	// By default this is the := value for all values
	var numArray [5]int
	var numString [5]string

	fmt.Println("default int is: ", numArray)
	fmt.Println("default for string is: ", numString)

	// We can also initially set the value of an array
	instantiatedArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(instantiatedArray)

	// We can also bring back the double array from 132
	var goodOle [5][5]int
	for i := 0; i < len(goodOle); i++ {
		for j := 0; j < len(goodOle[i]); j++ {
			goodOle[i][j] = i + j
		}
	}
	fmt.Println(goodOle)

}
