// Pset4.3 rewrite reverse to use an array pointer instead of a slice.
package main

import "fmt"

func reverse(s *[5]int) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Println(s) // "[5 4 3 2 1]"
}
