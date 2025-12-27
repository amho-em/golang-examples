package main

import "fmt"

// Zero Values: (0+0i)

var cxe = -9i // complex128

func main() {
	num := 23 + 12i // complex128
	fmt.Printf("%v: %T\n", cxe, cxe)
	fmt.Printf("%v: %T\n", num, num)

	var (
		a complex64  // the set of all complex numbers with float32 real and imaginary parts
		b complex128 // the set of all complex numbers with float64 real and imaginary parts
	)
	fmt.Printf("%v: %T\n", a, a)
	fmt.Printf("%v: %T\n", b, b)
}
