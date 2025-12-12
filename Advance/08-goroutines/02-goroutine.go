package main

import (
	"fmt"
	"time"
)

func printval(x int) {
	fmt.Printf("value 2 %d \n", x)
}

func main() {
	go func(x int) {
		fmt.Printf("value %d \n", x)
	}(10)

	go printval(20)

	time.Sleep(time.Second)
}