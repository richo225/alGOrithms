package main

import (
	"math/rand"
	"testing"
)

func BenchmarkCountingSort(b *testing.B) {
	numItems := rand.Intn(10000) + 1
	max := rand.Intn(10000) + 1
	slice := makeRandomCustomerSlice(numItems, max)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countingSort(slice, max)
	}
	b.StopTimer()

	// Print some statistics about the benchmark.
	b.Logf("Ran %d iterations of countingSort", b.N)
	b.Logf("Slice size: %d", numItems)
	b.Logf("Max value: %d", max)
}
