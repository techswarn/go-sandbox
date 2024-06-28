package main

import (
	"fmt"
	"example.com/greetings"
)

func main(){
	message := greetings.hello("Swarn")
	fmt.Println(message)
}