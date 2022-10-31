package main

import (
	"fmt"
	"os"
)

func echo_first(){
	s, sep := "",""
	fmt.Println(os.Args[1])
	for _, value := range os.Args[1:] {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
}

