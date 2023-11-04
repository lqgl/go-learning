// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Input: $ go run main.go 1 2 3 4
// Output: 
// 1 2 3 4 

// Different declare methods Same effects:
/*
	s := ""
	var s string
	var s = ""
	var s string = ""
*/
