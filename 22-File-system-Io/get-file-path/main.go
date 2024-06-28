package main

import (
    "os"
    "fmt"
    "path/filepath"
)

func main() {
    wd, err := os.Getwd()
    if err != nil {
        panic(err) 
    }
    fmt.Println(wd)
    content, err := os.ReadFile(filepath.Join(wd, "config", "config.json"))
    if err != nil {
        panic(err) 
    }
    fmt.Println(content)
}