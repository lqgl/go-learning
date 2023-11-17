// Pset3.10 a iteration version.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123"))       // "123"
	fmt.Println(comma("12345"))     // "12,345"
	fmt.Println(comma("123456789")) // "123,456,789"
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var buf bytes.Buffer
	res := len(s) % 3
	if res == 0 {
		res = 3
	}
	buf.WriteString(s[:res])
	s = s[res:]
	for len(s)/3 > 0 {
		buf.WriteByte(',')
		buf.WriteString(s[:3])
		s = s[3:]
	}
	return buf.String()
}
