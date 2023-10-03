package main

import (
	"fmt"
	"time"
	"runtime"
)


func start() {
	println("Executed start function")
}

func checkcores(){
	fmt.Println(runtime.NumCPU())
}

func main() {
	checkcores()
	fmt.Println("Start")
	go start()
	fmt.Println("Stop")
	time.Sleep(3 * time.Second)

}

