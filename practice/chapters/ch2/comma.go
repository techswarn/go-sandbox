package main

import "fmt"

// 1 2 3 4 5
// 0 1 2 3 4 5 6
func comma(s string) string {
	n := len(s);
	if n < 3 {
		return s
	}
	fmt.Println(s[:n-3])

	return comma(s[:n-3]) + "," + s[n-3:]
}