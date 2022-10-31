package main

import (
	"bytes"
	"fmt"
	// "os"
)

func intsToString( values []int) string {
	var b bytes.Buffer;
	// b.Write([]byte("Hello"))
	// fmt.Fprintf(&b, " World!")
	// b.WriteTo(os.Stdout)
	b.WriteByte('[')
	for i, v := range values {
		fmt.Fprintln(&b)
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d", v)
	}
	b.WriteByte(']')
	return b.String()
}