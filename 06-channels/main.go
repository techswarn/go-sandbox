package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	
	links := []string{"http://127.0.0.1:3000/api/v1/blogs"}
	c := make(chan string)

	for _, link := range links {
		for i:=0; i<1000; i++ {
			go statusChecker(link, c)
		}
	}
	for l := range c {
		fmt.Println(l)
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