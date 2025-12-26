package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	DisplayHelloWorld(os.Stdout)
}

func DisplayHelloWorld(w io.Writer) {
	fmt.Fprintln(w, "Hello, World!")
}
