package main

import (
	"math/rand"
	"testing"

	helpers "github.com/richo225/alGOrithms"
)

func BenchmarkQuickSort(b *testing.B) {
	numItems := rand.Intn(10000) + 1
	max := rand.Intn(10000) + 1
	slice := helpers.MakeRandomSlice(numItems, max)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSort(slice)
	}
	b.StopTimer()

	// Print some statistics about the benchmark.
	b.Logf("Ran %d iterations of quickSort", b.N)
	b.Logf("Slice size: %d", numItems)
	b.Logf("Max value: %d", max)
}
