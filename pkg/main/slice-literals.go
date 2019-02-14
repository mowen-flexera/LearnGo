package main

import "fmt"

func main() {
	var q = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	var r = []bool{true, false, true, true, false, true}
	fmt.Println(r)

	var s = []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s[1])
}
