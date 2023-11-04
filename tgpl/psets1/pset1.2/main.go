// Pset1.2 prints index and value of its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args[1:] {
		fmt.Printf("index: %v, value: %v\n", i, val)
	}
}

// Input: $: go run main.go 3 5 7 6

/* Output:
index: 0, value: 3
index: 1, value: 5
index: 2, value: 7
index: 3, value: 6
*/