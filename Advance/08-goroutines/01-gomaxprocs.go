package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("you are using ", runtime.Compiler, " ")
	fmt.Println("on a ", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version)
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}