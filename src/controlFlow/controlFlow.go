package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic if else
	x := 2
	if x+x == 5 {
		fmt.Println("math is wrong")
	} else if x+x == 6 {
		fmt.Println(" is wrong")
	} else {
		fmt.Println("2 + 2 is 4 quick maths")
	}

	// We can set variables in the if statement
	// But ONLY one is allowed
	if y := 9; y < 10 {
		fmt.Println(y, " is less than 10")
	} else if y > 10 {
		fmt.Println(y, "is more than 10")
	} else {
		fmt.Println(y, " is equal to 10")
	}

	// Variables in the conditional statement only exist
	// in the scope of the conditional block. The following
	// code would no compile.
	// fmt.Println(y)

	// Basic for loop from Java.
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// We can also do indefinite fors
	// With a conditional break
	i := 0
	for {
		fmt.Print(i, " ")
		i++
		if i == 10 {
			fmt.Println()
			break
		}
	}

	// There is no while loop in go. We have to use
	// the above for loop. But we can use a boolean
	i = 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// Also has my boi the switch statement.
	hour := time.Now().Hour()
	switch hour {
	// I don't like how defualt syntax is to not pad these.
	case 24:
		fmt.Println("it is midnight")

	case 12:
		fmt.Print("high noon")

	// We can use commas for multiple
	case 1, 13:
		fmt.Println("it is 1 am or pm")

	default:
		fmt.Println("the hour shows nothing of significance")
	}

}
