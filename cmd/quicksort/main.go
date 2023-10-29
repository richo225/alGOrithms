package main

import (
	"fmt"

	helpers "github.com/richo225/alGOrithms"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := helpers.MakeRandomSlice(numItems, max)
	helpers.PrintSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	quickSort(slice)
	helpers.PrintSlice(slice, 40)

	// Verify that it's sorted.
	helpers.CheckSorted(slice)
}

func quickSort(slice []int) {
	// check base case ie. array length is 0 or 1
	if len(slice) < 2 {
		return
	}

	// Partition the slice
	pivotIndex := partition(slice)

	// Recursively sort high and low halves
	quickSort(slice[0:pivotIndex])
	quickSort(slice[pivotIndex+1:])
}

func partition(slice []int) int {
	// Set the lower and upper indexes.
	lo := 0
	hi := len(slice) - 1

	// Use the last element as the pivot.
	pivot := slice[hi]

	// Temporary pivot index
	i := lo - 1

	for j := lo; j < hi; j++ {
		// See if slice[j] <= pivot.
		if slice[j] <= pivot {
			// Move the temporary pivot index forward
			i = i + 1

			// Swap slice[i] and slice[j].
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	// Drop the pivot between teh two halves.
	i = i + 1
	slice[i], slice[hi] = slice[hi], slice[i]

	// Return the pivot's index.
	return i
}
