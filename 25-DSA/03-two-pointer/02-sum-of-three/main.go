package main

	
import (
    "fmt"
	"sort"
//	"strings"
)

func findSumOfThree(nums []int, target int) bool {
	fmt.Println(nums)
	sort.Sort(sort.IntSlice(nums)) 
	fmt.Println(nums)

	low, high, triple := 0,0,0

	for i:=0; i < len(nums) - 2; i++ {
		low = i+1
		high = len(nums) - 1

		for low < high {
			fmt.Printf("Low value: %d \n",nums[low])
			fmt.Printf("High value: %d \n", nums[high])
			triple = nums[i] + nums[low] + nums[high]
			fmt.Printf("triple: %d \n", triple)
			fmt.Printf("Target: %d \n", target)
			if triple == target {
				return true
			} else if triple < target {
				low += 1
			} else if triple > target {
				high -=1
			}
	        
		}

	}
	return false
}

func main() {
	nums := []int{-1,2,1,-4,5,-3}
	target := 8
	res := findSumOfThree(nums, target)
    fmt.Printf("Result: %t \n", res)
}