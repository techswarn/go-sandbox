package main

import "fmt"

func months() {
	months := []string{"jan", "feb", "march", "april", "may", "june", "july",
						"august", "sep", "oct", "Nov", "dec"}
	nums := []int{1,2,3,4,5,6,7,8,9,10}

	for i,j:= 0, len(nums)-1; i<j ; i,j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)
	for _, m :=range months {
		fmt.Println(m)
	}

	
}