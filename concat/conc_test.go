package main

import (
	"testing"
)

func BenchmarkStandart(b *testing.B) {
	str := []string{"afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn"}
	for i := 0; i < b.N; i++ {
		_ = concat(str)
	}
}

func BenchmarkBuffer(b *testing.B) {
	str := []string{"afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn", "afa;dslkfn"}
	for i := 0; i < b.N; i++ {
		_ = concatBuffer(str)
	}
}
