package main

import (
	"math/rand"
	"testing"

	helpers "github.com/richo225/alGOrithms"
)

func BenchmarkBinarySearch(b *testing.B) {
	numItems := rand.Intn(10000) + 1
	max := rand.Intn(10000) + 1
	slice := helpers.MakeRandomSlice(numItems, max)
	targetIndex := rand.Intn(numItems)
	target := slice[targetIndex]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		binarySearch(slice, target)
	}
	b.StopTimer()

	// Print some statistics about the benchmark.
	b.Logf("Ran %d iterations of binarySearch", b.N)
	b.Logf("Slice size: %d", numItems)
	b.Logf("Max value: %d", max)
	b.Logf("Target index: %d", targetIndex)
}
