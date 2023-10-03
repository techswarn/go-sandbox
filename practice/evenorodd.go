package main

import (
	"fmt"
)

func main() {
	//Take a num
	var n int;
	fmt.Println("Enter the number: ")
	fmt.Scan(&n)
	//Divide by 2 to check if even, check reminder
	r := n%2
	if r == 0 {
		fmt.Printf("%d is even \n", n)
	} else {
		fmt.Printf("%d is odd \n", n)
	}
}