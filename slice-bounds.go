package main

import "fmt"

func main() {
	var s = []int{2, 3, 5, 7, 9, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
