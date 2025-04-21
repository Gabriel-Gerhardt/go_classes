package main

import "fmt"

type automovel interface {
	buzinar()
}

type Car struct {
	tamanho, combustivel int
}
type Moto struct {
	tamanho, combustivel int
}

func main() {
	car := Car{2, 4}
	moto := Moto{23, 4}

	car.buzinar()
	moto.buzinar()

}

func (m Moto) buzinar() {
	fmt.Println("vrumm")
}

func (c Car) buzinar() {
	fmt.Println("beep beep")

}
