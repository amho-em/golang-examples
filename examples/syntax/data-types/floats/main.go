package main

import "fmt"

// Zero Values: 0

var point1 = 16.75 // float64

func main() {
	point2 := -2.9871 // float64
	fmt.Printf("%v: %T\n", point1, point1)
	fmt.Printf("%v: %T\n", point2, point2)

	var (
		a float32 // the set of all IEEE 754 32-bit floating-point numbers
		b float64 // the set of all IEEE 754 64-bit floating-point numbers
	)
	fmt.Printf("%v: %T\n", a, a)
	fmt.Printf("%v: %T\n", b, b)
}
