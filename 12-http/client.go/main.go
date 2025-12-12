package main

import (
	"fmt"
	"net/http"
	//"encoding/json"
	"time"
)

func main() {
	fetch("https://jsonplaceholder.typicode.com/todos/1")
}

func fetch(url string) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response : %s", resp)
}