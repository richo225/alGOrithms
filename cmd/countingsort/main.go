package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Customer struct {
	id           string
	numPurchases int
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	slice := makeRandomCustomerSlice(numItems, max)
	printCustomerSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	sorted := countingSort(slice, max)
	printCustomerSlice(sorted, 40)

	// Verify that it's sorted.
	checkSorted(sorted)
}

func countingSort(slice []Customer, max int) []Customer {
	// Create counts slice of length max+1 all with entry 0
	counts := make([]int, max+1)
	// Iterate through slice and update count
	for _, v := range slice {
		counts[v.numPurchases] += 1
	}

	// Convert counts into the number of items less than or equal to each value
	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	// Make a new slice to hold the sorted results
	sortedSlice := make([]Customer, len(slice))

	// Loop through the original items from back to front.
	for i := len(slice) - 1; i >= 0; i-- {
		// Find value in modified count and insert into new slice at index = (value - 1)
		valueInCounts := counts[slice[i].numPurchases]
		sortedSlice[valueInCounts-1] = slice[i]
		// subtract 1 from counts
		counts[slice[i].numPurchases] -= 1
	}

	return sortedSlice
}

func makeRandomCustomerSlice(numItems int, max int) []Customer {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	slice := make([]Customer, numItems)
	for i := 0; i < numItems; i++ {
		slice[i].id = "C" + strconv.Itoa(i)
		slice[i].numPurchases = random.Intn(max)
	}
	return slice
}

func printCustomerSlice(slice []Customer, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []Customer) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1].numPurchases > slice[i].numPurchases {
			return false
		}
	}
	return true
}
