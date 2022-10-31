package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"https://www.google.co.in/", "https://www.spacex.com/"}
	c := make(chan string)

	for _, link := range links {
		go statusChecker(link, c)
	}

	for l := range c {
		go func (l string){
			time.Sleep(5 * time.Second)
			statusChecker(l, c)
		}(l)
	}
}

func statusChecker(link string, c chan string) {
		
		resp, err := http.Get(link)
		if err != nil {
			fmt.Println(link, " is down")
			c <- link
			return
		}
		fmt.Println("The status of website:", link, "is", resp.StatusCode)
		c <- link

}