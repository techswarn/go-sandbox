package main

import (
	"fmt"
	"net/http"
    "time"
	"sync"
	_ "bytes"
)

func main() {
	
	links := []string{"https://heap-app-q2axe.ondigitalocean.app/heapdump", "https://heap-app-q2axe.ondigitalocean.app/heapdump", "https://heap-app-q2axe.ondigitalocean.app/heapdump", "https://heap-app-q2axe.ondigitalocean.app/heapdump", "https://heap-app-q2axe.ondigitalocean.app/heapdump", "https://heap-app-q2axe.ondigitalocean.app/heapdump"}
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
		//buf := new(bytes.Buffer)
		resp, err := http.Get(link)
		//fmt.Printf("Response is: %v\n", resp)
		responseTime := time.Since(currentTime).Milliseconds()
		if err != nil {
			fmt.Println(link, " is down")
			c <- link
			return
		}
		//fmt.Printf("The status of website: %s is %d at time %d-%d-%d %d:%d:%d \n", link, resp.StatusCode, responseTime, currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second() )
        fmt.Printf("The status of website: %s is %d at time %d \n", link, resp.StatusCode, responseTime)
		c <- link

}
