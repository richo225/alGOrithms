package main

import (
	"fmt"
	"sort"
	"strconv"

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

	// Ask the user for a target
	for {
		var targetString string
		fmt.Printf("Target: ")
		fmt.Scanln(&targetString)

		// Exit the loop if target is empty.
		if targetString == "exit" || len(targetString) == 0 {
			break
		}

		target, err := strconv.Atoi(targetString)
		if err != nil {
			fmt.Println("Invalid input, please enter an integer.")
			continue
		}

		// Sort the slice
		// Search for target and display the result.
		index, numIterations := binarySearch(slice, target)
		if index < 0 || index >= len(slice) {
			fmt.Printf("Target %d not found, %d iterations\n", target, numIterations)
		} else {
			fmt.Printf("values[%d] = %d, %d iterations\n", index, target, numIterations)
		}
	}
}

func binarySearch(slice []int, target int) (int, int) {
	// Sort slice
	sort.Ints(slice)

	// Initialize left and right indices
	left := 0
	right := len(slice) - 1

	var numIterations int
	// Loop until left index is greater than right index
	for left <= right {
		numIterations += 1

		// Find the middle index
		middle := (left + right) / 2

		// If the middle element is the target, return its index
		if slice[middle] == target {
			return middle, numIterations
		}

		// If the target is less than the middle element, search the left half
		if target < slice[middle] {
			right = middle - 1
		} else { // Otherwise, search the right half
			left = middle + 1
		}
	}

	// If the target is not found, return -1 and numIterations
	return -1, numIterations
}
