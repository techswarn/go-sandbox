package main

import (
	"fmt"
	"unicode"
)

func main() {
	testCases := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"1A@2!3 23!2@a1",
		"No 'x' in Nixon",
		"12321",
	}

	for _, string := range testCases {
		fmt.Println(isPalindrome(string))
	}
}

func isPalindrome(s string) bool {
	fmt.Printf("Checking if string %s is a palindrome \n", s)
	//Set two pointers
	left, right := 0, len(s)-1

	for left < right {
		// fmt.Println(s[left])
		// fmt.Println(s[right])
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])){
			//Check if it's a number or alphabet
			left ++
		}

		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])){
			//Check if it's a number or alphabet
			right ++
		}
		
        //Now check if the char are palindone
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		left ++
		right --
	}
	return true
}