package main

import (
	"fmt"
)

func main() {
	var num[100] int
	var sum, avg, temp int
	fmt.Println("Enter the number: ")
	fmt.Scanln(&temp)
	for i:=0; i<=temp; i++{
		fmt.Scanln(&num[i])
		sum += num[i]
	}

	avg = sum/temp
	fmt.Printf("Average of %d number(s) is %d", temp, avg)
}