// Explanation
// The binary search algorithm is used to find a specific value in the sorted list. At each step, we look at the middle index. If the item at the middle index is the same as the desired one, 
// it is returned. Otherwise, the middle value is larger or smaller than the value we’re looking for. If our
//  desired value is smaller, we confine our search space to the left half of the array and ignore the right half 
//  and vice versa. We prune half of the search space at each stage. Because of that,this algorithm is more efficient than the linear search algorithm.

package main

import "fmt"

func main() {
 arr := []int{1,2,3,4,5,6,7,8,9,10,11,12,14}
 fmt.Println(arr[1:3])
 fmt.Println(BinarySearch(arr, 30))
}

func BinarySearch(data []int, value int) bool { 
    //Implement your solution here
    l := len(data)
    mid := l/2
    //check if middle index is the number
    if value == data[mid] {
        return true
    } else if value < data[mid] {
        for _, v := range data[0:mid] {
            if v == value {
                return true
            }
        }
    } else {
        for _, v := range data[mid:] {
            if v == value {
                return true
            }
        }
    }
    return false     //Return false if value not found
}

//ALTERNATIVE
// func BinarySearch(data []int, value int) bool { 
//     var mid int
//     low := 0
//     high := len(data) - 1

//     for low <= high {
//         mid = (low + high)/2
//         if data[mid] == value { // values find
//             return true
//         } else {
//             if data[mid] < value { // move to left
//                 low = mid + 1
//             } else {
//                 high = mid - 1   // move to right
//             }
//         }
//     }
//   return false                  // value not found 
// }