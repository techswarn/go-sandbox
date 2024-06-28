package main

import (
	"fmt"
	"sync"
)

// sum is used to calculate some sum of added numbers.
type sum struct {
	sum int
	mu sync.Mutex
}

// get returns current value of the sum.

func (s *sum) get() int {
	// This locks so that only one get() or add() call
	// can acces sum.sum at a time. This gives us the
	// required syncronization.
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.sum
}

// add adds 'n' to the sum value.
func (s *sum) add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sum += n
}

func main() {
	mySum := &sum{}
	// wg is used to let us know when all goroutines are done.
	wg := sync.WaitGroup{}
	for i:=0; i<100; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			mySum.add(x)
		}(i)
	}
	wg.Wait()
	fmt.Println("final sum", mySum.get())
}