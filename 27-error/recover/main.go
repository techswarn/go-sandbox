package main
import (
  "fmt"
)

func badCall() {
// panic("bad end")
fmt.Println("function did not panic")
}

func test() {
  defer func() {
    if e := recover(); e != nil {
      fmt.Printf("Panicking %s\r\n", e);
    }
  }()
  badCall()
  fmt.Printf("After bad call\r\n");
}

func main() {
  fmt.Printf("Calling test\r\n");
  test()
  fmt.Printf("Test completed\r\n");
}