package main

import (
	"fmt"
)

func sumArr(data []int) int { 

	total := 0
	  
	//Implement your solution here
	for _ , val := range data {
		total = total + val
	}
  
	return total
}

func main() {
	fmt.Println(sumArr([]int{1,2,3,4,5,6}))
}

