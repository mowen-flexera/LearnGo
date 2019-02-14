package main

import "fmt"

func main() {
	var s = []int{2, 3, 5, 7, 9, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[1:5]
	printSlice(s)

	// Drop its first wto values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
