package main

import "fmt"

type Stack struct {
	stack []int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
}

func (s *Stack) Pop() int {
	if len(s.stack) == 0 {
		fmt.Println("Error: Stack is empty")
		return 0
	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *Stack) Size() int {
	return len(s.stack)
}

