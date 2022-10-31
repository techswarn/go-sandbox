package main

import "fmt"

type bot interface {
	getGreetings() string
}
type engbot struct {}
type spanbot struct {}

func main() {
	eb := engbot{}
	sb := spanbot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreetings())
}

func (en engbot) getGreetings() string {
	return "hi there"
}

func (sp spanbot) getGreetings() string {
	return "holla amego"
}