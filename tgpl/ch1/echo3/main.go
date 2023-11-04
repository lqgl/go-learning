// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Input: $ go run main.go 1 2 3 4
// Output: 
// 1 2 3 4 
