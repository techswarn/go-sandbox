package main

import "fmt"

func main() {
	fmt.Println(SequentialSearch([]int{1,2,3,4,5,30}, 30))
}

func SequentialSearch(data []int, value int) bool { 
    //Implement your solution here
	for _, val := range data {
		if value == val {
			return true
		} 
	}
    return false     //Return false if value not found
}