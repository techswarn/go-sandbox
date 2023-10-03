package main

import (
	"fmt"
)

func mul(n int) {

	for i:=1;i <=10; i++ {
		num := n * i
		fmt.Println(num)
	}


}

func main() {
	fmt.Println("Enter the number:")
    var n int
	fmt.Scan(&n)
	fmt.Printf("Generating multiplications for %d \n", n)
	mul(n)
}