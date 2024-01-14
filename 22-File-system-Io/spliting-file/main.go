package main

import (
    "os"
    "fmt"
    "path/filepath"
    "io"
)

func main() {
    fp := "/usercode/config/config.json"
    fileName := filepath.Base(fp)
    if fileName == "." {
        fmt.Println("Path is empty")
    }

    newPath := filepath.Join(os.TempDir(), fileName)

    r, err := os.Open(fp)
    if err != nil {
        panic(err)
    }
    defer r.Close()

    w, err := os.OpenFile(newPath, os.O_WRONLY | os.O_CREATE, 0644)
    if err != nil { 
        panic(err)
    }
    defer w.Close()
    // Copies the file to the temporary file.
    _, err = io.Copy(w, r)
    if err != nil { 
        panic(err)
    }
}