// Pset3.11 reports whether two strings are anagrams of each other.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(anagrams("abc", "dcba"))            // "false"
	fmt.Println(anagrams("abcde", "decba"))         // "false"
	fmt.Println(anagrams("abcd", "dcba"))           // "true"
	fmt.Println(anagrams("世界", "世界"))               // "false"
	fmt.Println(anagrams("Hello, 世界", "界世 ,olleH")) // "true"
}

func anagrams(a, b string) bool {
	aC := utf8.RuneCountInString(a)
	bC := utf8.RuneCountInString(b)
	if aC != bC {
		return false
	}
	aRunes := []rune(a)
	bRunes := []rune(b)
	for i := 0; i < aC; i++ {
		j := aC - 1 - i
		if aRunes[i] != bRunes[j] {
			return false
		}
	}
	return true
}
