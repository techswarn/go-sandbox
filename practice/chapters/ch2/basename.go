package main

import (
	"fmt"
	"strings"
)

// func basename(s string) string {
// 	//Discard / and everything before
// 	//a/b/cd
// 	for i := len(s) - 1; i >=0; i--{
// 		fmt.Println(s[i])
// 		if s[i] == '/' {
// 			s = s[i+1:]
// 			break
// 		}
// 	}

// 	//preserve everything before last '.'
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == '.'{
// 			s = s[:i]
// 			break
// 		}
// 	}
// 	return s
// }

func basename (s string) string {
	slash := strings.LastIndex(s, "/")
	fmt.Println(slash)
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
		fmt.Println(s)
	}
	return s
}