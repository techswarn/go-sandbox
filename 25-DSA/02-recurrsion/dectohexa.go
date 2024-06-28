package main

import (
	"fmt"
)

/*

95%16 = 5 | 15 > F
5%16 = 0 | 5

Explanation
Line 5: We initialize a conversion variable that we’ll use to convert a base-10 digit to base-16.
Line 6: We define a base variable to use later in the code.
Line 7: We take the modulus of number using the base variable to separate a base-16 digit from the number and save it to the digit variable.
Line 8: We divide the number with base to make changes in the number variable too.
Line 9–11: We first check if number is not equal to 0. If true, the printInt(number) is called recursively. If false, the function will not be called recursively and works as a base case.
Line 12: After exiting the if condition, the digit variable is used as an index of the conversion variable to finally convert the base 10 number to the base 16 number. Finally, we print the converted digit to the console.
*/

func main() {
	fmt.Println("decimal to hexa conversion")
	dechexacon(95)

}

func dechexacon(n int) {
	conversion := "0123456789ABCDEF"
	base := 16
	r := n % base
	q := n / base
	if( n != 0) {
		dechexacon(q)
	}
	
	fmt.Print(string(conversion[r]))
}