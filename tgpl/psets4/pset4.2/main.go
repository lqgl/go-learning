// Pset4.2 prints the SHA256/SHA384/SHA512 hash of stdin.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var width = flag.Int("w", 256, "hash width with '384 or 512', default: 256")

func main() {
	flag.Parse()
	var sha func([]byte) []byte
	switch *width {
	case 256:
		sha = func(b []byte) []byte {
			hash := sha256.Sum256(b)
			return hash[:]
		}
	case 384:
		sha = func(b []byte) []byte {
			hash := sha512.Sum384(b)
			return hash[:]
		}
	case 512:
		sha = func(b []byte) []byte {
			hash := sha512.Sum512(b)
			return hash[:]
		}
	default:
		log.Fatal("unexpected with specified.")
		os.Exit(1)
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read from stdin: %v\n", err)
		os.Exit(2)
	}
	fmt.Println(sha(b))
}
