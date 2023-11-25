package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c) // "5", = len("Hello")

	c = 0 // reset
	var name = "Dolly"
	fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c) // "12", = len("Hello, Dolly")
}
