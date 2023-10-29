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

	// select a pivot
	pivot := (slice)[len(slice)-1]

	var less []int
	var greater []int
	for i := 0; i < len(slice)-1; i++ {
		if (slice)[i] <= pivot {
			// set slice of all values less than pivot
			less = append(less, (slice)[i])
		} else {
			// set slice of all values greater than pivot
			greater = append(greater, (slice)[i])
		}
	}

	// run quicksort on less
	quickSort(less)

	// run quicksort on greater
	quickSort(greater)

	// combine together and update original slice
	copy(slice, append(append(less, pivot), greater...))
}
