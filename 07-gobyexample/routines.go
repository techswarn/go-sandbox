package main

import (
	"fmt"
	// "time"
)

func f (from string) {
	for i:=0; i<3; i++ {
		fmt.Println(from,";", i)
	}
}


func routines() {
	f("hello")

	go f("routine")

	go func(msg string) {
        fmt.Println(msg)
    }("going")

	// time.Sleep(time.Second)
    fmt.Println("done")
}