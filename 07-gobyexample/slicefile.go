package main

import "fmt"

func slicefile(){
	s := make([]string, 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	fmt.Println("set :", s)
	fmt.Println("get :", s)
	fmt.Println("Length: ", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied new slice: ", c)
	l := s[2:5]
	fmt.Println("slice with new set of elements ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        // innerLen := i + 1
        twoD[i] = make([]int, 3)
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}