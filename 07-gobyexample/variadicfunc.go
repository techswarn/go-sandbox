package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums)
	total :=0

	for _, num := range nums {
		total = total + num
	}
	return total
}

func variadicfunc() {
	fmt.Println(sum(10, 20, 30))
}