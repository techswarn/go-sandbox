//array = { 8, 1, 2, 3, 4, 5, 6, 4, 2 }
//array = [ 8, 1, 3, 2, 5, 4, 6, 2, 4 ] //output

package main

import ("fmt")

func main() {
	arr := []int{8, 1, 2, 3, 4, 5, 6, 4, 2}

	size := len(arr)

	for i:=1 ; i<size; i+=2 {
		fmt.Println(arr[i])
		if (i-1) >= 0 && arr[i] > arr[i-1] {
			swap(arr, i, i-1)
		} 
		if (i+1) < size && arr[i] > arr[i+1] {
			swap(arr, i, i+1)
		}
		fmt.Println(arr)
	}
	fmt.Println(arr)
}


func swap(data []int, k int, j int) {
		data[k], data[j] = data[j], data[k]	
}