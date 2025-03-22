package main

import (
	"fmt"
)

// insertEnd inserts n into arr at the next open position.
func insertEnd(arr []int, n, length, capacity int) []int {
	if length < capacity {
		arr[length] = n
		length++
	}
	return arr
}

// removeEnd removes from the last position in the array if the array
// is not empty (i.e., length is non-zero).
func removeEnd(arr []int, length int) []int {
	if length > 0 {
		arr[length-1] = 0
		length--
	}
	return arr
}

// insertMiddle inserts n into index i after shifting elements to the right.
// Assuming i is a valid index and arr is not full.
func insertMiddle(arr []int, i, n, length int) []int {
	// Shift starting from the end to i.
	for index := length - 1; index > i-1; index-- {
		arr[index+1] = arr[index]
	}
	// Insert at i
	arr[i] = n
	length++
	return arr
}

// removeMiddle removes value at index i before shifting elements to the left.
// Assuming i is a valid index.
func removeMiddle(arr []int, i, length int) []int {
	// Shift starting from i + 1 to end.
	for index := i + 1; index < length; index++ {
		arr[index-1] = arr[index]
	}
	// No need to 'remove' arr[i], since we already shifted
	arr[length-1] = 0
	length--
	return arr
}

// printArr prints the array up to the length
func printArr(arr []int, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
