package main

import "fmt"

type Square struct {
	side int
}

func main() {
	square := Square{20}

	fmt.Println(square.Area())
	fmt.Println(area(square))
}

func (s Square) Area() int {
	return s.side * s.side
}

func area(s Square) int {
	return s.side * s.side
}
