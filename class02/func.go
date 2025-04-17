package main

import (
	"fmt"
)

func main() {
	value, value2 := 3, 3
	fmt.Printf("%v\n", pow(&value, &value2))
}

func pow(x *int, y *int) int {
	value := *x

	for i := 1; i < *y; i++ {
		value = *x * value
	}
	return value
}
