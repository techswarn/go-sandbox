package main

import (
	"fmt"
)

func main() {
	var rows int
	fmt.Println("Enter the number of pyramid rows: ")
	fmt.Scan(&rows)

	for i:=1; i<=rows; i++ {
		c:=rows*2-1

		for j:=1; j<=c; j++ {

				fmt.Print("*  ")
		}
		fmt.Println(" ")
	}
}