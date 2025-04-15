package main

import (
	"fmt"
)

func main() {
	value := 2000
	transformToByte(&value)
	fmt.Printf("%b\n", value)
}

func transformToByte(i *int) {
	*i = int(uint8(*i))
}
