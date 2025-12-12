package main

import "fmt"

func main() {
	var arr [5]int
	arr[4] = 10

	for i,v := range arr {
		fmt.Println(i,v)
	}
}