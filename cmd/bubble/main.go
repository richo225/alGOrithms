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

	// Make and display an unsorted slice.
	slice := helpers.MakeRandomSlice(numItems, max)
	helpers.PrintSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	helpers.PrintSlice(slice, 40)

	// Verify that it's sorted.
	fmt.Println(helpers.CheckSorted(slice))
}

func bubbleSort(slice []int) []int {
	if helpers.CheckSorted(slice) {
		return slice
	}
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
	}
	return bubbleSort(slice)
}
