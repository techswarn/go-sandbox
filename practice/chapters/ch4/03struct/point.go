package main

import "fmt"

type Point struct{x,y int}

func Scale(p Point, factor int) Point {
	return Point{ p.x * factor , p.y * factor}
}

func Pointfunc(){
	fmt.Println(Scale(Point{1,2}, 5))
}
