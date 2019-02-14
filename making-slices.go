package main

import "fmt"

func main() {
	var a = make([]int, 5)
	printSlice("a", a)

	var b = make([]int, 0, 5)
	printSlice("b", b)

	var c = b[:2]
	printSlice("c", c)

	var d = c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
