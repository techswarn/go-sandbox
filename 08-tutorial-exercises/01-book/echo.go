package main

import (
	"fmt"
	"os"
)

func echo(){
	var s, sep string
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
