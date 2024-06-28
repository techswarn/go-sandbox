//Sample input
// array = { 1,2,3,4,5,6 }
// rotatedarry = { 3,4,5,6,1,2 }
// k = 2

package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5,6}
	fmt.Println(RotateArray(arr, 3))
}

func RotateArray(data []int, k int) []int {
	size := len(data)
	fmt.Println(ReverseArray(data, 0, k-1))
	fmt.Println(ReverseArray(data, k, size-1))
	fmt.Println(ReverseArray(data, 0, size-1))
	

	return data
    //You must make changes in the given array
}

func ReverseArray(data []int, start int, end int) []int {
	i := start
	j := end

	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
	return data
}