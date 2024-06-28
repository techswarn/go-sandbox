package main

import "fmt"

func rangefile() {
	nums := []int{1,2,3,4,5,6}
	sum :=0
	for _, num :=range nums {
		sum = sum + num
	}
	fmt.Println(sum)

	for i, value := range nums {

		if(value == 3){
			fmt.Println("The index with value: ", value, "is ", i)
		}	
	}

	kvs := map[string]string{"a": "apple", "b": "ball", "c": "cat"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for i, c := range "go" {
        fmt.Println(i, c)
    }
}