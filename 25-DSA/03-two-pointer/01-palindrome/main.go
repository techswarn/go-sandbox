package main

import "fmt"

func isPalindrome(inputstr string) bool{

	fmt.Printf("Checking if %s is a plaindrome \n ", inputstr)

	left := 0
	right := len(inputstr) - 1
	for left < right {
		if inputstr[left] != inputstr[right] {
			return false
		}
		left +=1
		right -=1
	}
	return true
}

func main() {
	words := []string {"hello", "racecar"}
	for _, val := range words{
		res := isPalindrome(val)
		fmt.Printf("%s is a panlindrome: %t \n", val, res)
	}

}



// func twoPointers(array []int) {
// 	left := 0
// 	right := len(array) - 1
// 	for left < right {
// 		left = left + 1
// 		right = right - 1
// 	}
// }