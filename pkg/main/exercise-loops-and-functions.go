package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	fmt.Println("Original")
	z := float64(1)

	for i := 0; i < 10; i++ {

		z -= (z*z - x) / (2 * z)
		//fmt.Println(z)
	}
	return z
}

func Sqrt_flow_control(x float64) float64 {
	fmt.Println("Flow Controlled")
	z := float64(x / 2)
	z_og := float64(x/2 - 1)
	//fmt.Println(z, z_og)
	z_diff := float64(1.0)
	//fmt.Println(z_diff)
	for z_diff > .000000001 {
		z_og = z
		z -= (z*z - x) / (2 * z)
		z_diff = z - z_og
		if z_diff < 0 {
			z_diff = z_diff * -1
		}
		//fmt.Println(z_diff)
	}
	return z
}

func main() {
	ex := 42.0
	fmt.Println("Answer:", Sqrt(ex))
	fmt.Println("Answer:", Sqrt_flow_control(ex))
}
