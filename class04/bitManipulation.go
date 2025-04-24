/*Binário no pelo, usa 0b para binário(padrão em todas linguagens) */
package main

import (
	"fmt"
)

func main() {
	var b uint8 = 0b01001010
	var a uint8 = 0b10010010
	fmt.Println(b & a) //and
	fmt.Println(b | a) //or
}
