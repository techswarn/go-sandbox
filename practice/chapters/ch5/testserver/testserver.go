package testserver

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) (error) {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	fmt.Println(deadline)
	
	for tries :=0; time.Now().Before(deadline); tries++ {
		res, err := http.Get(url)
		fmt.Println(res.StatusCode)
		if err == nil {
			return nil
		}
		log.Printf("Server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}