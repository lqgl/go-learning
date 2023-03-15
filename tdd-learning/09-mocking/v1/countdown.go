package main

import (
	"bytes"
	"fmt"
)

func Countdown(buffer *bytes.Buffer) {
	fmt.Fprintf(buffer, "3")
}
