package main

import (
	"fmt"
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

		// Search for target and display the result.
		index, numIterations := linearSearch(slice, target)
		if index < 0 || index >= len(slice) {
			fmt.Printf("Target %d not found, %d iterations\n", target, numIterations)
		} else {
			fmt.Printf("values[%d] = %d, %d iterations\n", index, target, numIterations)
		}
	}
}

func linearSearch(slice []int, target int) (int, int) {
	// Iterate through slice
	for i, v := range slice {
		// If target is found, return index and numIterations
		if v == target {
			return i, i + 1
		}
	}
	// If target not found, return index -1 and numIterations
	return -1, len(slice)
}
