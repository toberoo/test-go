package main

import (
	"fmt"
)

// Handle value parameters
type colorful interface {
	getColor() string
}

type car struct {
	paintColor string
}

type food struct {
	color string
}

func (c car) getColor() string {
	return c.paintColor
}

func (f food) getColor() string {
	return f.color
}

func printColor(c colorful) {
	fmt.Println(c.getColor())
}

// Handle pointer handlers
type animal interface {
	greet() string
}

type dog struct{}
type cat struct{}

func (*dog) greet() string {
	return "bark bark"
}

func (*cat) greet() string {
	return "meow meow"
}

func petAnimal(a animal) {
	fmt.Println(a.greet())
}

func main() {

	// Let's try the value one
	redCar := car{"red"}
	blueberry := food{"blue"}
	printColor(redCar)
	printColor(blueberry)

	// Need to grab the refrence as only
	// the refrence has the parameters. This
	// is stupid confusing
	d := &dog{}
	c := &cat{}
	petAnimal(d)
	petAnimal(c)
}
