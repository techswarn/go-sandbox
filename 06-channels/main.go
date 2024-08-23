package main

import (
	"fmt"
	"net/http"
    "time"
	"sync"
)

func main() {
	
	links := []string{"https://imgprocess-app-qeopc.ondigitalocean.app/process"}
	c := make(chan string)
	wg := sync.WaitGroup{}
	for _, link := range links {
			go statusChecker(link, c)
	}
	for l := range c {
		fmt.Println(l)
		wg.Add(1)
		go func (l string){
			defer wg.Done()
			time.Sleep(4 * time.Second)
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
