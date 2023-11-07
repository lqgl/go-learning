// Pset2.4 Benchmark shift version and table lookup version.
package main

import "testing"

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(1000000)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(1000000)
	}
}

/*
$: go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/lqgl/go-learning/tgpl/psets2/pset2.4
BenchmarkPopCount1-8    1000000000               0.3190 ns/op
BenchmarkPopCount2-8    56218048                21.27 ns/op
PASS
ok      github.com/lqgl/go-learning/tgpl/psets2/pset2.4 2.034s
*/