package main

import "fmt"

type Grupo struct {
	pessoas []*Pessoa
}
type Pessoa struct {
	name string
}

func main() {
	grupo := Grupo{}
	pessoa := Pessoa{name: "john"}
	grupo.addPessoa(pessoa)
	fmt.Println(grupo.pessoas[0].name)
}

func (g *Grupo) addPessoa(pessoa Pessoa) {
	g.pessoas = append(g.pessoas, &pessoa)

}
