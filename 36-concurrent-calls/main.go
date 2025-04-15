package main

import (
    "fmt"
    "sync"
	"net/http"
	"log"
)

func worker(id int) {
   // fmt.Printf("Worker %d starting\n", id)

	resp, err := http.Get("https://lb.techenv.dev/")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("RESPONSE status code: %d \n", resp.StatusCode)
   // fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 1000; i++ {
        wg.Add(1)

        go func() {
            defer wg.Done()
            worker(i)
        }()
    }

    wg.Wait()

}