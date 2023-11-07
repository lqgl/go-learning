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
pkg: github.com/lqgl/go-learning/tgpl/psets2/pset2.3
BenchmarkPopCount1-8    1000000000               0.3175 ns/op
BenchmarkPopCount2-8    319338507                3.758 ns/op
PASS
ok      github.com/lqgl/go-learning/tgpl/psets2/pset2.3 2.166s
*/