package main

import (
	"log"
	"net/http"
	"fmt"
	"sync"
)

type Site struct {
	URL string
}

type Result struct {
	URL, workerIdMsg string
	Status int
}

func pingWebsite(wId int, jobs <-chan Site, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
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
}

func main(){
	var wg sync.WaitGroup
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w:=0; w <= 1000;w++ {
		wg.Add(1)
		go pingWebsite(w, jobs, results, &wg)
	} 

	urls := []string{}
	for i :=0; i<=9999; i++ {
		urls = append(urls, "https://lb.techenv.dev/")
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