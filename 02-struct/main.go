package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	//First way
	swarn := person{firstName: "Swarn", lastName: "Suvarna"}
	fmt.Println(swarn)
	// fmt.Printf("%+v", swarn)
	// swarn.firstName = "John"
	// fmt.Printf("%+v", swarn)
	
	//second way
	// var swarn person
	// fmt.Println(swarn)
	// fmt.Printf("%+v", swarn)

	jim := person{
		firstName: "jim",
		lastName: "doe",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipcode: 345556,
		},
	}
	//way 1
	// jimPointer  := &jim
	jim.updateName("Sam")
	jim.print()
}

func (pointerToPerson *person)updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}