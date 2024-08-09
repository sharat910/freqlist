package main

import "testing"

func BenchmarkFreqList(b *testing.B) {
	words := getWords()
	for i := 0; i < b.N; i++ {
		topUsingFreqList(words, 10, 10)
	}
}

func BenchmarkFreqMap(b *testing.B) {
	words := getWords()
	for i := 0; i < b.N; i++ {
		topExact(words, 10)
	}
}
