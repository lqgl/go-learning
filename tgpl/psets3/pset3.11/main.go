// Pset3.11 enhance comma to deals correctly with floating-point numbers and an optional sign.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("+132.23"))          // "+132.23"
	fmt.Println(comma("+132"))             // "+132"
	fmt.Println(comma("+1331412.233141"))  // "+1,331,412.233141"
	fmt.Println(comma("-13231412.233141")) // "-13,231,412.233141"
}

func comma(s string) string {
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	// Deal with floating-point numbers.
	eI := strings.Index(s, ".")
	if eI == -1 {
		eI = len(s)
	}
	r := len(s[:eI]) % 3
	if r > 0 {
		buf.WriteString(s[:r])
		buf.WriteByte(',')
	}

	for i, v := range s[r:eI] {
		if i%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(v)
	}

	buf.WriteString(s[eI:])
	return buf.String()
}
