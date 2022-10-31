package main

import "fmt"

func pointers() {
	var a int = 10;
	fmt.Println(a)
	fmt.Println(&a)
	var p *int
	p = &a
	fmt.Println(*p)
}