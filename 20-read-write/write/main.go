package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	fmt.Println("Writing to file: ")
	outputFile, outputError := os.OpenFile("output/output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
	  fmt.Printf("An error occurred with file creation\n")
	  return
	}
	defer outputFile.Close()
	outputWriter:= bufio.NewWriter(outputFile)
	outputString := "hello world!\n"
	for i:=0; i<10; i++ {
	  outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}

//FOR SIMPLER TASK
// package main
// import "os"

// func main() {
//   os.Stdout.WriteString("hello, world\n")
//   f, _ := os.OpenFile("output/test.txt", os.O_CREATE|os.O_WRONLY, 0)
//   defer f.Close()

//   f.WriteString("hello, world in a file\n")
// }