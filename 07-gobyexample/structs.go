package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name string, age int) *person {
	p1 := person{name, age}
	return &p1
}

func (p *person) updateName(name string) {
	p.name = name
}

func structs() {
	p := person{name:"swarn", age:29}
	fmt.Println(p)
	p.updateName("sam")
	fmt.Println(p)

	person2 := newPerson("Rooke", 45)
	fmt.Println(*person2)
}


