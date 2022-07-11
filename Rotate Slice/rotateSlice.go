package main

import "fmt"

func main() {
	str := "PropertyFinder"
	var strSlice []string

	// Create a slice contains all letters of string value
	for i := 0; i < len(str); i++ {
		strSlice = append(strSlice, string(str[i]))
	}

	fmt.Println(strSlice)                          // [P r o p e r t y F i n d e r]
	fmt.Println(RotateSlice(strSlice, 2, "left"))  // [o p e r t y F i n d e r P r]
	fmt.Println(RotateSlice(strSlice, 2, "right")) // [e r P r o p e r t y F i n d]

}

// RotateSlice rotate the given slice by given n number to the right or to the left
func RotateSlice(strSlice []string, n int, direction string) []string {
	// Basically, divide given slice by two part with n number index
	// After that combine these parts each other with append function
	var result []string
	if direction == "left" {
		s1 := strSlice[n:]
		s2 := strSlice[:n]
		result = append(s1, s2...)
	} else if direction == "right" {
		s1 := strSlice[len(strSlice)-n:]
		s2 := strSlice[:len(strSlice)-n]
		result = append(s1, s2...)
	}
	return result
}
