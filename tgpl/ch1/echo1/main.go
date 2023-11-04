// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// os.Args[0] is the command name invoked
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Input: $ go run main.go 1 2 3 4
// Output: 
// 1 2 3 4 

/* for loop
for initialization; condition; post {
	// zero or more statements
}
*/