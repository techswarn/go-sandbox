package main

import (
	"bufio"
	_ "encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title string
	Price float64
	Quantity int
}

func main() {

	bks := make([]Book, 1)
	inputFile, inputError := os.Open("product.txt")
	if inputError != nil {
		log.Fatal(inputError)
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	for {
		// read one line from the file:
		line, err := inputReader.ReadString('\n')
		if err == io.EOF {
			break 
		}

//		fmt.Printf("Line: %s \n", line)
		str := strings.Split(line, ";")
//		fmt.Printf("line 1: %s \n", str)

		book := new(Book)
		book.Title = str[0]
		book.Price, err = strconv.ParseFloat(str[1], 32)
		if err != nil {
			log.Fatal(err)
		}
		book.Quantity, err = strconv.Atoi(str[2])
		fmt.Printf("%v", book)

		if bks[0].Title == "" {
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}