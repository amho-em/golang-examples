package main

import "fmt"

// Zero Value: "" (an empty string)

var name = "Amirhossein Emadi" // string

func main() {
	city := "Gorgan" // string
	fmt.Printf("%v: %T\n", name, name)
	fmt.Printf("%v: %T\n", city, city)

	var a string
	fmt.Printf("%q: %T\n", a, a)
}
