package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi " + name + ", How are you")
	return message
}
