package main

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	numItems := rand.Intn(10000) + 1
	max := rand.Intn(10000) + 1
	slice := makeRandomSlice(numItems, max)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bubbleSort(slice)
	}
	b.StopTimer()

	// Print some statistics about the benchmark.
	b.Logf("Ran %d iterations of bubbleSort", b.N)
	b.Logf("Slice size: %d", numItems)
	b.Logf("Max value: %d", max)
}
