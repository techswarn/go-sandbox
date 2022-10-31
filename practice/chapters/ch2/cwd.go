package main

import (
	"log"
	"os"
)

var cwd string

func cwdfunc() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Current working directory: %v", cwd)
}