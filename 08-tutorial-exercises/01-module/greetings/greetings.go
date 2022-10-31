package greetings

import (
	"errors"
	"fmt"
)

func hello(name string) string{

	if name == "" {
		errors.New("empty name")
		return
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}