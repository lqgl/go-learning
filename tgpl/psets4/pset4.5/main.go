// Pset4.5 writes an in-place function to eliminate adjacent duplicates in a []string slice.
package main

import "fmt"

func duplicate(strings []string) []string {
	i := 0
	for _, s := range strings {
		if strings[i] == s {
			continue
		}
		i++
		strings[i] = s
	}
	return strings[:i+1]
}

func main() {
	var s = []string{"a", "b", "b", "c"}
	fmt.Println(duplicate(s)) // "["a" "b" "c"]"
}
