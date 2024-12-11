package main

import (
	"fmt"
	"net/http"
    "time"
	"sync"
)

func main() {
	
	links := []string{"https://go.techenv.dev/api/v1/", "https://go.techenv.dev/api/v1/", "https://go.techenv.dev/api/v1/", "https://go.techenv.dev/api/v1/", "https://go.techenv.dev/api/v1/"}
	c := make(chan string)
	wg := sync.WaitGroup{}
	for _, link := range links {
			go statusChecker(link, c)
	}
	for l := range c {
		wg.Add(1)
		go func (l string){
			defer wg.Done()
			statusChecker(l, c)
		}(l)
	}
}

func statusChecker(link string, c chan string) {
		currentTime := time.Now()
		_ = currentTime
		resp, err := http.Get(link)
		if err != nil {
			fmt.Println(link, " is down")
			c <- link
			return
		}
		fmt.Printf("The status of website: %s is %d at time %d-%d-%d %d:%d:%d \n", link, resp.StatusCode, currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second() )
		c <- link

}
