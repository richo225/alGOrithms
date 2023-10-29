package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	fmt.Println(checkSorted(slice))
}

func bubbleSort(slice []int) []int {
	if checkSorted(slice) {
		return slice
	}
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			slice[i-1], slice[i] = slice[i], slice[i-1]
		}
	}
	return bubbleSort(slice)
}

func makeRandomSlice(numItems int, max int) []int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		slice[i] = random.Intn(max)
	}
	return slice
}

func checkSorted(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}
