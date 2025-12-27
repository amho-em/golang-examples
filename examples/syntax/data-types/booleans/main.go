package main

import "fmt"

/*
-> Zero Values: false
-> Values: 'true' (1) or 'false' (0)
*/

var isMale = true // bool

func main() {
	isChecked := false // bool
	fmt.Printf("%v: %T\n", isMale, isMale)
	fmt.Printf("%v: %T\n", isChecked, isChecked)

	var a bool
	fmt.Printf("%v: %T\n", a, a)
}
