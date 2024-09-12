package main

import (
	"fmt"
    "time"
	"sync"
	"net"
	"os"
)

func main() {
	
	links := []string{"google.com"}
	c := make(chan string)
	wg := sync.WaitGroup{}
	for _, link := range links {
			go dnsStatusChecker(link, c)
	}
	for l := range c {
		fmt.Println(l)
		wg.Add(1)
		go func (l string){
			defer wg.Done()
			time.Sleep(10 * time.Second)
			dnsStatusChecker(l, c)
		}(l)
	}
}

func dnsStatusChecker(link string, c chan string) {
		ips, err := net.LookupIP(link)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
			c <- link
			return
		}
		for _, ip := range ips {
			fmt.Printf("google.com. IN A %s\n", ip.String())
			c <- link
		}
}