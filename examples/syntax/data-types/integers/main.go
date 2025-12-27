package main

import "fmt"

/*
-> Zero Values: 0
-> u: Unsigned Integers (only positive numbers)
*/

var number = -16 // int

func main() {
	age := 25 // int
	fmt.Printf("%v: %T\n", age, age)
	fmt.Printf("%v: %T\n", number, number)

	var (
		a int8  // -128 through 127
		b uint8 // 0 through 255

		c int16  // -32,768 through 32,767
		d uint16 // 0 through 65,535

		e int32  // -2,147,483,648 through 2,147,483,647
		f uint32 // 0 through 4,294,967,295

		g int64  // -9,223,372,036,854,775,808 through 9,223,372,036,854,775,807
		h uint64 // 0 through 18,446,744,073,709,551,615
	)
	fmt.Printf("%v: %T\n", a, a)
	fmt.Printf("%v: %T\n", b, b)
	fmt.Printf("%v: %T\n", c, c)
	fmt.Printf("%v: %T\n", d, d)
	fmt.Printf("%v: %T\n", e, e)
	fmt.Printf("%v: %T\n", f, f)
	fmt.Printf("%v: %T\n", g, g)
	fmt.Printf("%v: %T\n", h, h)

	var (
		i int  // 'int32' or 'int64' (depends on your OS)
		t uint // 'int32' or 'int64' (depends on your OS)

		r rune // alias for 'int32' (character values)

		o byte // alias for 'uint8' (byte values)
	)
	fmt.Printf("%v: %T\n", i, i)
	fmt.Printf("%v: %T\n", t, t)
	fmt.Printf("%v: %T\n", r, r)
	fmt.Printf("%v: %T\n", o, o)
}
