package main

import "testing"

// max_test.go
func BenchmarkMaxInt(b *testing.B) {
	sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxInt(sl)
	}
}

func BenchmarkMaxAny(b *testing.B) {
	sl := []any{1, 2, 3, 4, 7, 8, 9, 0}
	for i := 0; i < b.N; i++ {
		maxAny(sl)
	}
}

/*
$go test -v -bench . ./max_test.go max_any.go max_int.go
goos: darwin
goarch: amd64
... ...
BenchmarkMaxInt
BenchmarkMaxInt-8     398996863           2.982 ns/op
BenchmarkMaxAny
BenchmarkMaxAny-8     85883875          13.91 ns/op
PASS
ok    command-line-arguments  2.710s
*/
