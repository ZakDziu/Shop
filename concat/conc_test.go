package main

import (
	"strings"
	"testing"
)

const (
	oneTestData = "this is a test string"
	testCount   = 10000
)

var (
	expected = strings.Repeat(oneTestData, testCount)
	source   []string
)

func init() {
	for i := 0; i < testCount; i++ {
		source = append(source, oneTestData)
	}
}

func BenchmarkConcat(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concat(source)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}

func BenchmarkConcatOptimized(b *testing.B) {
	var result string
	for n := 0; n < b.N; n++ {
		result = concatOptimized(source)
	}
	b.StopTimer()
	if result != expected {
		b.Errorf("unexpected result; got=%s, want=%s", result, expected)
	}
}
