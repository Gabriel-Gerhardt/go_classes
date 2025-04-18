package main

import (
	"fmt"
)

type Car struct {
	model string
	name  string
}
type BigCar struct {
	Car
	size int
}

func main() {

	c := Car{"Lamborghini", "Aventador"}
	b := BigCar{c, 20}
	fmt.Println(b.getSize())
}

func (c BigCar) getSize() int {
	return c.size
}
