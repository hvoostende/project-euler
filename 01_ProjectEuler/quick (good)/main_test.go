package main

import "testing"

func BenchmarkProgram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		totalNumbers(i)
	}
}
