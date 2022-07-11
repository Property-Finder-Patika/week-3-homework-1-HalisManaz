package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice []string
	str := "PropertyFinder"

	// Create a slice contains all letters of string value
	strSlice = strings.Split(str, "")

	fmt.Println(strSlice)                          // [P r o p e r t y F i n d e r]
	fmt.Println(RotateSlice(strSlice, 2, "left"))  // [o p e r t y F i n d e r P r]
	fmt.Println(RotateSlice(strSlice, 2, "right")) // [e r P r o p e r t y F i n d]

}

// RotateSlice rotate the given slice by given n number to the right or to the left
func RotateSlice(strSlice []string, n int, direction string) []string {
	// Basically, divide given slice by two part with n number index
	// After that combine these parts each other with append function
	var s1, s2 []string
	if direction == "left" {
		s1 = strSlice[n:]
		s2 = strSlice[:n]
	} else if direction == "right" {
		s1 = strSlice[len(strSlice)-n:]
		s2 = strSlice[:len(strSlice)-n]
	}
	return append(s1, s2...)
}
