package main

import (
	"bytes"
	"fmt"
)

func main() {

	// Creating and initializing strings
	// Using bytes.Buffer with
	// WriteString() function
	var b bytes.Buffer

	b.WriteString("M")
	b.WriteString("a")
	b.WriteString("v")
	b.WriteString("e")
	b.WriteString("n")
	b.WriteString("i")
	b.WriteString("r")

	fmt.Println("String: ", b.String())

	b.WriteString(" ")
	b.WriteString("G")
	b.WriteString("o")
	b.WriteString("L")
	b.WriteString("a")
	b.WriteString("n")
	b.WriteString("g")
	b.WriteString(" ")

	fmt.Println("String: ", b.String())

}
