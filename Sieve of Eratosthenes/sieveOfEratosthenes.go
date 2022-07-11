package main

import (
	. "HW3/functions"
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	n := 100
	result := SliceUpToN(n)
	fmt.Println(result)

	primeNumbers := Eratosthenes(100)
	fmt.Println(primeNumbers)
}

func Eratosthenes(n int) []int {
	// Create empty slice for storing prime numbers
	var primeSlice []int
	// Create a slice from 2 to n
	mySlice := SliceUpToN(n)

	// Try all numbers from 2 to n
	for i := 2; i <= n; i++ {
		// Check the value in slice or not
		if slices.Contains(mySlice, i) {
			primeSlice = append(primeSlice, i)
			if i*i <= n {
				for j := i * i; j <= n; j = j + i {
					// Check that the value to be removed is not already removed
					if slices.Contains(mySlice, j) {
						// Find index desired value
						index := FindIndex(mySlice, j)
						// Remove the value from slice
						mySlice[index] = mySlice[len(mySlice)-1]
						mySlice = mySlice[:len(mySlice)-1]
					} else {
						// If it is already removed continue
						continue
					}
				}
			} else {
				// If i*i greater than n -> continue
				continue
			}
		} else {
			// if the value is not in slice continue
			continue
		}
	}
	return primeSlice
}
