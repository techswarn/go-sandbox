package main

import (
	"fmt"
)

func main() {
	var num, temp, rem int
	var result int = 0
	fmt.Println("Enter the number: ")
	fmt.Scan(&num)

	temp = num

	for {
		rem = num % 10
		result += rem * rem * rem
		num = num / 10
		if num == 0 {
			break
		}
	}

	if result == temp {
		fmt.Printf("%d is armtrong number \n", temp)
	} else {
		fmt.Printf("%d is not armtrong number \n", temp)
	}
}
