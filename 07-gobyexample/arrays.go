package main

import "fmt"
func arrays(){
	var a [5]int;
	fmt.Println("emp: ", a);

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("Get:", a[4])
	fmt.Println("Length of array is: ", len(a))

	b := [6]int{0,1,2,3,4,5}
	fmt.Println(b)

	var twoD [2][3]int
	fmt.Println(twoD)
	for i:=0; i < 2; i++ {
		for j:=0; j<3; j++ {
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d array is ", twoD)
}