package main

import "fmt"

func main() {
	var arr [15]int;
	for i:=0; i < len(arr) ; i++ {
		arr[i] = i
	}
	fmt.Println(fibs())
}

// fibonacci series using array

var fib[10]int64

func fibs() [10]int64{

	var a, b int64 = 0 , 1

	for i:=0; i<len(fib); i++ {
		if i < 1 {
			fib[i] = 1
			continue
		}
		a ,b = b , a+b
		fib[i] = b
	}

    return fib
}

