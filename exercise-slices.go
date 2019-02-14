package main

import "pic"

func Pic(dx, dy int) [][]uint8 {

	mypic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		mypic[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			mypic[i][j] = uint8(((1/2)*i - j) ^ i - j)
		}
	}
	return mypic
}

func main() {
	pic.Show(Pic)
}
