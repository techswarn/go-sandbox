package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Testing go concurrency");
	go spinner(100 * time.Millisecond)
	log.Println(fib(45))
}

//Calculate fibonacci number

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x - 1) + fib(x - 2)
}

//Create a spinner
func spinner(delay time.Duration) {
	for {
		for _, r := range `- \|/` {
			log.Println("\r%c", r)
			time.Sleep(delay)
		}
	}
}