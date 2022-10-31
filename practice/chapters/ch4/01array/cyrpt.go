package main

import (
	"fmt"
	"crypto/sha256"
)

func cyrpt(){
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}