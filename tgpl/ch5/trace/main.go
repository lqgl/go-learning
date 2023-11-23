package main

import (
	"log"
	"time"
)

func BigSlowOperation() {
	defer trace("bigSlowOperation...")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	BigSlowOperation()
}

// Output:
// 2023/11/23 15:57:05 enter bigSlowOperation...
// 2023/11/23 15:57:15 exit bigSlowOperation... (10.0014985s)
