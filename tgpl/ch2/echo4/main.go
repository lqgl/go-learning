// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // update defualt flag variable value
	fmt.Print(strings.Join(flag.Args(), *sep)) // sep: *string
	if !*n { // n: *bool
		fmt.Println()
	}
}

/*
$: go build -o echo4 main.go
$: ./echo4 a bc def
a bc def
$: ./echo4 -s a bc def
a bc def
$: ./echo4 -s / a bc def
a/bc/def
$: ./echo4 -n a bc def
a bc def$
$: ./echo4 -help
Usage of ./echo4:
  -n    omit trailing newline
  -s string
        separator (default " ")
*/