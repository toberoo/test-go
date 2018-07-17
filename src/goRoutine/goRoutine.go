package main

import (
	"fmt"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printLetters() {
	for _, c := range "Iknowwhatyoudid" {
		fmt.Println(string(c))
	}
}

func main() {

	// The print will be intermingled
	go printNumbers()
	go printLetters()

	fmt.Scanln()
	fmt.Println("done")
}
