package main

import "fmt"

func main() {
	arr := []int{ 8, 5, 6, 1, 9, 11, 2, 7, 4, 10 }
    size := len(arr)
	fmt.Println(checksmallsol1(arr, size))
}

func checksmallsol1(data []int, size int) int {

	// h := 0

	// for _, val := range data {
	// 	if val > h{
	// 		h = val
	// 	}
		
	// }
     
	// for pass:=0; pass < size; pass ++ {
	// 	for i:=pass+1; i<size ;i++ {
	// 		if data[pass] > data[i] {
	// 			data[pass], data[i] = data[i], data[pass]
	// 		}
	// 	}
	// }
	// fmt.Println(data)
	// for _, val := range data {
	// 	if val > h{
	// 		h = val
	// 	}
		
	// }

	found := 0

	for j:=1 ; j < size +1 ; j++ {
		found = 0
		for i:=0; i< size; i++ {
			if data[i] == j {
				found = 1
			}
		}
		if found == 0 {
			return j
		}
	}

	return -1
}


func checksmallsol2(data []int, size int) int {

	hs := make(map[int]int)
	for i := 0; i < size; i++ {
		  hs[data[i]] = 1
	}
	  for i := 1; i < size+1; i++ {
		  _, ok := hs[i]
		  if ok == false {
			  return i
		  } 
	  }
  
  return -1 
}

func checksmallsol3(data []int, size int) int {

	hs := make(map[int]int)
	for i := 0; i < size; i++ {
		  hs[data[i]] = 1
	}
	  for i := 1; i < size+1; i++ {
		  _, ok := hs[i]
		  if ok == false {
			  return i
		  } 
	  }
  
  return -1 
}
