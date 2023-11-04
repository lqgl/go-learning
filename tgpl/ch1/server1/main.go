// Serve1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// $: go run main.go &
// [1] 90521
// $: ps -a
/*
  PID TTY           TIME CMD
45518 ttys002    0:00.53 /bin/zsh -il
90521 ttys002    0:00.24 go run main.go
90545 ttys002    0:00.02 /var/folders/vm/3mhy12_9037_dmx1mw_nql9c0000gn/T/go-build277252031/b001/exe/main
90553 ttys002    0:00.01 ps -a
*/

// Stop background running program
// $: kill 90545