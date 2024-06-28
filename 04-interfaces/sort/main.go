package main

import (
	"fmt"
)

type Person struct {
	firstname string
	lastname string
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons )Less(i, j int) bool {
	pi := p[i].firstname +  " " + p[j].lastname
	pj := p[j].firstname + " " + p[j].lastname

	return pi < pj
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i] 
}

func main() {
	fmt.Println("Invoked")

	p1 := Person{"Swarn", "suvarna"}
	p2 := Person{"Mark", "lure"}
	p3 := Person{"Dickson", "Desouza"}

	persons := Persons{p1, p2, p3}
	fmt.Println(persons)
}