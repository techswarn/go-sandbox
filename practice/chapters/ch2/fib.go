package main

import "fmt"

func fib(n int) int {
	x, y := 0, 1
	for i:=0; i<n; i++{
		x,y = y,x+y
		fmt.Printf("x vaiable: %d \n", x)
		fmt.Printf("y variable: %d \n", y)
	}
	fmt.Printf("x before return: %d \n", x)
	return x
}