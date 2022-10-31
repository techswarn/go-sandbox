package main

import "fmt"

const boilingF = 212.0

func boiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g c\n", f, c)
}