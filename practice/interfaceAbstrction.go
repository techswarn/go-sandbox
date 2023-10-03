package main

import (
	"fmt"
)

type Triangle struct {
	height, base float64
}

type Rectangle struct {
	length, breath float64
}

// Method for type triangle
func (t Triangle) Area() float64{
	return 0.5* t.height * t.base
}

func (r Rectangle) Area() float64{
	return r.length * r.breath
}

type Area interface {
	Area() float64
}

func main() {
	t := Triangle{height: 10, base: 16}
	r := Rectangle{length: 10, breath: 20}
	var a Area

	a = t
	fmt.Println("Area of triangle is: ", a.Area())
	a = r 
	fmt.Println("Area of rectangle: ", a.Area())

}

