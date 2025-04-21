package main

import "fmt"

type Numeros struct {
	values  []int
	values2 []*int
}

func main() {
	value := 3
	num := Numeros{values: []int{1, 2}, values2: []*int{&value}}

	fmt.Println(*num.values2[0])
}
