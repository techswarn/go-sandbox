package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	limit := 0
	for input.Scan(){
		counts[input.Text()]++
		limit ++
		if limit == 10 {
			break
		}
	}
	fmt.Println(counts)

	for line, n := range counts {
		if n >1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}