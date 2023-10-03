package main

import (
	"net/http"
	"fmt"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "Hello world", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}