// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // "["one" "three"]"
	fmt.Printf("%q\n", data)           // "["one" "three" "three"]"
	data2 := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty2(data2)) // "["one" "three"]"
	fmt.Printf("%q\n", data2)            // "["one" "three" "three"]"
	s := []int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))
}
