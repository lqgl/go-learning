package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const url = "https://www.bad.gopl.io"

// WaitForServer attempts to contact the server of a URL.
// It tries for one minute using exponential back-off.
// It reports an error if all attempts fail.
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}

// Output:
/*
2023/11/22 15:17:22 server not responding (Head "https://www.bad.gopl.io": EOF); retrying...
2023/11/22 15:17:23 server not responding (Head "https://www.bad.gopl.io": EOF); retrying...
2023/11/22 15:17:25 server not responding (Head "https://www.bad.gopl.io": EOF); retrying...
2023/11/22 15:17:30 server not responding (Head "https://www.bad.gopl.io": EOF); retrying...
2023/11/22 15:17:38 server not responding (Head "https://www.bad.gopl.io": EOF); retrying...
2023/11/22 15:17:54 server not responding (Head "https://www.bad.gopl.io": EOF); retrying...
Site is down: server https://www.bad.gopl.io failed to respond after 1m0s
exit status 1
*/
