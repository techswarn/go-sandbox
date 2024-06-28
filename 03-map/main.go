package main

import "fmt"

func main(){
	colors := map[string]string{
		"white": "shirt",
		"blue": "pant",
		"brown": "shoe",
	}

	// var colors map[string]string
	// colors1 := make(map[string]string)

	colors["grey"] = "tie"

	// delete(color, "grey")

	fmt.Println(colors)
	printMap(colors)
}

func printMap (c map[string]string) {
	for color, resource := range c	{
		fmt.Println(color + " is for " + resource)
	}
}
