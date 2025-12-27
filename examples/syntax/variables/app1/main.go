package main

import "fmt"

var (
	name string
	age  int
)

func main() {
	name = "Amirhossein Emadi"
	age = 25
	fmt.Printf("My name is %s. I'm %d years old.\n", name, age)

	a := "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	d := true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
