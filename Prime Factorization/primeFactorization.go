package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"math"
)

func main() {

	factorization := make(map[int]int)
	n := 8
	mySlice := Eratosthenes(n)
	for _, num := range mySlice {
		count := 1
		for n%int(math.Pow(float64(num), float64(count))) == 0 {
			factorization[num] = count
			count++
		}

	}
	fmt.Printf("%+v", factorization)

	//fmt.Println(100 % int(math.Pow(2, 5)))
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

func FindIndex(arr []int, desiredValue int) int {
	// Find index value in slice for desired value
	for index, value := range arr {
		if value == desiredValue {
			return index
		}
	}
	return -1 // No found
}

func SliceUpToN(n int) []int {
	// Create a slice from 2 to N
	var mySlice []int
	for i := 2; i <= n; i++ {
		mySlice = append(mySlice, i)
	}

	return mySlice
}
