package main

import (
	"fmt"
	"os"
	"golang.org/x/net/html"
	"./testserver"
)

func main(){
//	findlinks()
	servercheck("https://google.com")
}

func findlinks () {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Println("Error parsing")
		os.Exit(1)
	}
	fmt.Println(doc)
}

func servercheck (url string){
	if err := testserver.WaitForServer("https://googles.com"); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}