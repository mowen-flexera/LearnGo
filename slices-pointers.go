package main

import "fmt"

func main() {
	var names = [4]string{
		"John",
		"Paul",
		"Van",
		"Damn",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
