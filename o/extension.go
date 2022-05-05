package main

import "fmt"

type SayMyName struct {
	name string
}

func (smn SayMyName) Say() { fmt.Printf("- Say my name!\n- %s!\n", smn.name) }

type DigaMeuNome struct {
	SayMyName
}

func (dmn DigaMeuNome) Say() { fmt.Printf("- Diga meu nome!\n- %s!\n", dmn.name) }

func main() {
	var myNameEn SayMyName
	myNameEn.name = "Heisenberg"
	var meuNomePt DigaMeuNome
	meuNomePt.name = "Heisenberg"

	fmt.Println("English version")
	myNameEn.Say()

	fmt.Println("")
	fmt.Println("Versão brasileira")
	meuNomePt.Say()
}
