package main

import "fmt"

// Zero Value: "" (an empty string)

var name = "Amirhossein Emadi" // string

func main() {
	city := "Gorgan\nTehran" // normal string (double quotes "...")
	info := `Amirhossein
Gorgan is a big city!

I live in Gorgan.` // raw string (backticks `...`)
	fmt.Printf("%v: %T\n", name, name)
	fmt.Printf("%v: %T\n", city, city)
	fmt.Printf("%v: %T\n", info, info)

	var a string
	fmt.Printf("%q: %T\n", a, a)
}
