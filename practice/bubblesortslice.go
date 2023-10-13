package main

import "fmt"

func main() {

	s := []int{4,5,2,1,3}
	sorted := bubbleSort(s)
	fmt.Println(sorted)
}

func bubbleSort(sl []int) []int{

	for pass :=0; pass < len(sl); pass++ {
		for i:= pass +1; i < len(sl); i++ {
			fmt.Println(sl[pass], sl[i])
			if sl[pass] > sl[i] {
				sl[pass], sl[i] = sl[i] ,sl[pass]
			}
		}
	}

	return sl
}