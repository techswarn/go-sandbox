package main

import "fmt"

type address struct {
	hostname string
	port int
}

func addressfunc() {
	hits := make(map[address]int)
	hits[address{"swarn.tech", 8080}]++
	hits[address{"swarn.tech", 8080}]++
	fmt.Println(hits)
}
