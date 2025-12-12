package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) == 2 {
		fmt.Println(os.Args[1])
	}
}

