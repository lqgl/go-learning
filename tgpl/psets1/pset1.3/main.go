// Pset1.3 measure the time difference between low efficency program and using strings.Join program.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// echo1
	start := time.Now()
	Echo1(os.Args)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	// echo2
	start = time.Now()
	Echo2(os.Args)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	// echo3
	start = time.Now()
	Echo3(os.Args)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func Echo1(args []string) {
	s, sep := "", ""
	for i := 1; i < len(args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}
