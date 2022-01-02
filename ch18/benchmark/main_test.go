package main

import "testing"

func BenchmarkStreamDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreamDecode("art.json")
	}
}

func BenchmarkMemoryDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MemoryDecode("art.json")
	}
}
