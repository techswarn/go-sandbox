package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

    for i := 0; i < checkCores(); i++ {
		wg.Add(1)
			
        go func(n int) {
			defer wg.Done();
			fmt.Println(n)
		}(i)// This happens concurrently
		
    }
	wg.Wait()
	fmt.Println("All work done")
}

func checkCores() int {
	cpuCores := runtime.NumCPU();
	fmt.Printf("Number of CPU cores: %d \n", cpuCores)
	return cpuCores
}
 