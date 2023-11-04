// Pset1.1 prints the name of the command invoked.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}

// Output: 
// /var/folders/vm/3mhy12_9037_dmx1mw_nql9c0000gn/T/go-build3745586461/b001/exe/main