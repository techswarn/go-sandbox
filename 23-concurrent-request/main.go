package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func pingWebsite(wid int, url string, wg *sync.WaitGroup) {
	log.Printf("Worker %d pinging website %s", wid, url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Worker %d: Error pinging website %s: %v", wid, url, err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Worker %d: Status code for %s: %d", wid, url, resp.StatusCode)

	defer wg.Done()
}

func createJobs(n int) {
	for i := 0; i < n; i++ {
		
	}
}

func main() {


	var wg sync.WaitGroup
	// get the number of concurrent workers and the jobs
	log.Printf("Workers: %s\n", os.Args[1])
	log.Printf("URL: %s \n", os.Args[2])
	log.Printf("Jobs: %d \n", len(os.Args[3]))
	url := os.Args[2]
	workers, err := strconv.Atoi(os.Args[1])
	jobs, err := strconv.Atoi(os.Args[3])
	if err != nil {	
		log.Println(err)
		return
	}

	go createJobs(jobs)

	for w:=0; w < workers; w++ {
		wg.Add(1)
		go pingWebsite(w, url, &wg)
	}
    defer wg.Wait()
	
}