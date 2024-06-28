package main

import "fmt"

func main() {
	fmt.Println("index array")
	arr := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	size := len(arr)
	indexArray(arr, size)
	fmt.Println(arr)
}

func indexArray(arr []int, size int) {
    //Implement your solution here
	fmt.Println(size)
    //Note: Please make changes in the given array

	//Sort the arr
	//  newarr := bubbleSort(arr)
	//  fmt.Println(newarr)
	for i:=0; i<size; i++ {
		curr := i
		value := -1
/* swaps to move elements in the proper position. */ 
		for arr[curr] != -1 &&  arr[curr] != curr {
			      
			temp := arr[curr]
			arr[curr] = value
			value = temp
			curr = temp
		}
/* check if some swaps happened. */
		if value != -1 {
            arr[curr] = value
        } 

	}
}

func bubbleSort(arr []int) []int {

	size := len(arr)

	for pass:=0; pass < size; pass++ {
		for i := pass+1; i < size; i ++ {
			if arr[pass] > arr[i] {
				arr[pass], arr[i] = arr[i], arr[pass]
			}
		}
	} 
	return arr
}