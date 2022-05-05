package main

import "fmt"

type SayMyName2 struct{}

func (smn SayMyName2) Name() string { return "I am Heisenberg!" }

func (smn SayMyName2) PrintName() {
	fmt.Println(smn.Name())
}

type DigaMeuNome2 struct {
	SayMyName2
}

func (dmn DigaMeuNome2) Name() string { return "Eu sou Heisenberg!" }

func main() {
	var heisenberg DigaMeuNome2
	fmt.Println(heisenberg.Name())
	heisenberg.PrintName()
}
