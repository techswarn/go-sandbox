package main
import (
	"fmt"
	"bufio"
	"os"
)

var (
      firstName, lastName, s string
      i int
      f float32
      input = "56.12 / 5212 / Go"
      format = "%f / %d / %s"
)
var inputReader *bufio.Reader
var input string
var err error


func main() {
  fmt.Println("Please enter your full name: ")
  fmt.Scanln(&firstName, &lastName)
  // fmt.Scanf("%s %s", &firstName, &lastName)
  fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
  fmt.Sscanf(input, format, &f, &i, &s)
  fmt.Println("From the string we read: ", f, i, s)

  inputReader = bufio.NewReader(os.Stdin)
  fmt.Println("Please enter some input: ")

  input, err = inputReader.ReadString('\n')
  if err == nil {
    fmt.Printf("The input was: %s\n", input)
  }
}