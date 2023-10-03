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

/* Practising inheritance */

type Position struct {
	x float32
	y float32
}

func(p *Position) setPosition(x,y float32) {
	x += x
	y += y
}

type Player struct {
	*Position
}

func NewPlayer() *Player{
	return &Player{
		Position: &Position{},
	}
}

func main(){
	//First way
	// var firstname, lastname string;
	// fmt.Println("Enter firstname and lastname:")
	// fmt.Scanln(&firstname , &lastname)
	// swarn := person{firstName: firstname, lastName: lastname}
	// fmt.Println(swarn)
	// // fmt.Printf("%+v", swarn)
	// // swarn.firstName = "John"
	// // fmt.Printf("%+v", swarn)
	
	// //second way
	// // var swarn person
	// // fmt.Println(swarn)
	// // fmt.Printf("%+v", swarn)

	// jim := person{
	// 	firstName: "jim",
	// 	lastName: "doe",
	// 	contact: contactInfo{
	// 		email: "jim@gmail.com",
	// 		zipcode: 345556,
	// 	},
	// }
	// //way 1
	// // jimPointer  := &jim
	// jim.updateName("Sam")
	// jim.print()

	splayer := NewPlayer()
	fmt.Println(splayer.Position)
}

func (pointerToPerson *person)updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}