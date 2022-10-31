package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func structsmethods() {
	box := rect{20, 30}
	fmt.Println(box)
	fmt.Println(box.area())

	copiedbox := &box
	fmt.Println(*copiedbox)
}