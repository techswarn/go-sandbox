package main

import (
	"fmt"
	"time"
)
func switchstatement(){
	i := 2;

	switch i {
	case 1:
		fmt.Println("The number set is :", i);
	case 2:
		fmt.Println("The number set is :", i)		
	}

	fmt.Println(time.Now().Weekday())
}