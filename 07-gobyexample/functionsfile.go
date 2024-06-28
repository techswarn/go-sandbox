package main

import "fmt"

func plus ( a int, b int) int {
	return a + b 
}

func plusPlus (a, b, c int) int {
	return a+b+c
}

func functionsfile(){
	fmt.Println(plus(2, 3))
	fmt.Println(plusPlus(5,6,7))
}