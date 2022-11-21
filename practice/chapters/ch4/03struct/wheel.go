package main

import "fmt"

type Wheel struct {
	Circle Circle
	Spokes int
}

type Circle struct {
	Center Point
	Radius int
}

type Point struct {
	X, Y int
}

var w = Wheel {
	Circle: Circle {
		 Center: Point {
			X : 5,
			Y : 10,
		},
		Radius : 20,
	},
	Spokes : 30,
}
func createWheel() {
	fmt.Printf("%#v\n", w)
}

