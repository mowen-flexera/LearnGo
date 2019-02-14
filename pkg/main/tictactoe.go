package main

import "fmt"

func displayframe() {
	defer fmt.Println("1|2|3\n4|5|6\n7|8|9")
	fmt.Println("_|_|_\n_|_|_\n | | ")
}

func input() int {
	fmt.Println("Enter location (1-9):")
	return 0
}

func main() {
	entry := 0
	displayframe()
	entry = input()
	fmt.Println(entry)
}
