package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9}
	res := minmax(arr, len(arr))
	fmt.Println(res)
}

func minmax(arr []int, size int) []int {
	aux := make([]int, size)
	copy(aux, arr)
	fmt.Println(aux)
	start := 0
	stop := size - 1
	for i:=0; i< size; i++ {
		if i%2==0 {
			arr[i] = aux[stop]
			stop --
		} else {
			arr[i] = aux[start]
			start++
		}
	}
	return arr
}

//Solution without using extra space
// package main

// import ("fmt")

// func swap(x, y int) (int, int) {
//     return y, x
// }

// func MaxMinArr2(arr []int, size int) { 
//   for i := 0; i < (size - 1); i++ {
//         ReverseArr(arr, i, size-1)
//     }
// }

// func ReverseArr(arr []int, start int, stop int) { 
//   for start < stop {
//         arr[start],arr[stop]=swap(arr[start], arr[stop])
//         start += 1
//         stop -= 1 
//   } 
// }
// /* Testing code */
// func main() {
//   arr := []int{1, 2, 3, 4, 5, 6, 7} 
//   size := len(arr)
//   MaxMinArr2(arr, size) 
//   fmt.Println(arr)
// }