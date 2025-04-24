package main

import (
	"fmt"
)

type Mail struct {
	message string
}

func main() {
	mail1 := Mail{"simmm"}
	mail2 := Mail{"valor"}
	arrays := []Mail{mail1, mail2}
	arr1, arr2 := splitArr(arrays)
	fmt.Println(arr1)
	fmt.Println(arr2)
}

/*Generics can have any name, but T is global convetion  */
func splitArr[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}
