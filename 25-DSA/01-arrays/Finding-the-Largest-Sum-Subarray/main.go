//array = { 1, -2, 3, 4, -4, 6, -14, 6, 2 }

package main

import "fmt"

func main() {
	arr := []int{ 1, -2, 3, 4, -4, 6, -14, 6, 2}
    fmt.Println(MaxSubArraySum(arr))
}

func MaxSubArraySum(data []int) int {
    l := len(data)
    //Implement your solution here
    maxSum := 0
    maxNum := 0
    for i:=0; i < l-1; i++ {
//        fmt.Println(data[i])
        maxNum += data[i]
        
        if maxNum < 0 {
            maxNum = 0
        } 

        if maxSum < maxNum  {
            maxSum = maxNum
        }
    }
    
    return maxSum
}