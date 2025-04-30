package main

import (
	"log"
	"net/http"
	"fmt"
	"sync"
	"os"
)

type Site struct {
	URL string
}

type Result struct {
	URL, workerIdMsg string
	Status int
}

func pingWebsite(wId int, jobs <-chan Site, results chan<- Result, wg *sync.WaitGroup) {
	for site := range jobs {
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		// sending into result channel
		results <- Result{
			workerIdMsg:	fmt.Sprintf("The worker id is %d, and status_code",wId),
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
	defer wg.Done()
}

func main(){
	var wg sync.WaitGroup

	fmt.Printf("OS ARGS: %s\n", os.Args)
	jobs := make(chan Site, 10)
	results := make(chan Result, 10)

	for w:=0; w <= 1000;w++ {
		wg.Add(1)
		go pingWebsite(w, jobs, results, &wg)
	} 

	urls := []string{}
	for i :=0; i<=9999; i++ {
		urls = append(urls, "https://spaces.techenv.dev/Screenshot%202025-03-10%20at%208.19.52%E2%80%AFPM.png")
	}

	    // sending into jobs channel
		for _, url := range urls {
			jobs <- Site{URL: url}
		}
		close(jobs)
	
		for i := 1; i <= len(urls); i++ {
			result := <-results
			fmt.Println(result)
		}
    wg.Wait()
}