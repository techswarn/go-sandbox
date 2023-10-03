package main

import (
	"log"
	"net/http"
	user "titan/user"

	"github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Gorilla!\n"))
}

func main() {
    r := mux.NewRouter()

    // Routes consist of a path and a handler function.
    r.HandleFunc("/", YourHandler)
	r.HandleFunc("/user", user.UserHanlder)

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}