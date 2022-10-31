package main

import "fmt"

func calculator(){
	fmt.Println("Welcome to calculator build using go")

	var operator string
	var num1, num2 int
	
	fmt.Println("Please enter the first number: ")
	fmt.Scanln(& num1)
	fmt.Println("Please enter the second number: ")
	fmt.Scanln(& num2)
	fmt.Println("What opearation would you like to perform: ")
	fmt.Scanln(& operator)
	result := 0
	switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
	}

	fmt.Println(result)
}