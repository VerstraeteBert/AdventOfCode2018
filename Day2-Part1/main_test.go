package main

import "testing"

func BenchmarkConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concurrent()
	}
}

func BenchmarkNonConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		nonConcurrent()
	}
}

