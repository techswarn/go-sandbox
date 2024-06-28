package main

import "fmt"

var i int = 5
var name string = "swarn"

type Person struct {
	name string
	age int
}

type any interface {}

func main(){
	var val any
	fmt.Printf("%#v \n", val)
//	val = i
	fmt.Printf("%#v \n", val)
	p := new(Person)
	p.age = 2
	p.name = "pikachu"
	fmt.Printf("%#v \n", p)

	switch t:= val.(type) {
	case int:
		fmt.Printf("%#v \n", t)
	case string:
		fmt.Printf("%#v \n", t)
	case *Person:
		fmt.Printf("%#v \n", p) 
	}
}