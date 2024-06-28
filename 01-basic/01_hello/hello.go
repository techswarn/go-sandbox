package main

import "fmt"

func main() {
	a, b, c, d, e, f := 0,2,3,5,6,7
	fmt.Println("Printing in binary")

	fmt.Printf("a is %b \n", a)
	fmt.Printf("b is %b \n", b)
	fmt.Printf("c is %b \n", c)
	fmt.Printf("d is %b \n", d)
	fmt.Printf("e is %b \n", e)
	fmt.Printf("f is %b \n", f)

	fmt.Println("Printing in hexadecimal")
	fmt.Printf("a is %x \n", a)
	fmt.Printf("b is %x \n", b)
	fmt.Printf("c is %x \n", c)
	fmt.Printf("d is %x \n", d)
	fmt.Printf("e is %x \n", e)
	fmt.Printf("f is %x \n", f)

}