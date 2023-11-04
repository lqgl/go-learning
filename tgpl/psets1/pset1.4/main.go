// Pset1.4 modify dup2 program, prints file name appears duplicate lines more than once.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "pset1.4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	tmpCounts := make(map[string]int)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		tmpCounts[line]++
	}
	// NOTE: ignoring potential errors from input.Err()

	// Print name of file that appears duplicate lines more than once.
	for _, n := range tmpCounts {
		if n > 1 {
			fmt.Printf("%s file appears duplicate lines more than once.\n", f.Name())
			return
		}
	}
}
