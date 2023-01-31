package main

import "fmt"

func main() {

	// Declaration using var key word

	// syntax var <variable> <data_type>

	var a int

	fmt.Printf("\n Type of a is %T and value of a is %d", a, a)

	var b = 0

	fmt.Printf("\n Type if b is %T and value is %d", b, b)

	var c, d = 0.0, "Hello world from go lang"

	fmt.Printf("\n Type of c is %T \n Type of d is %T", c, d)

	fmt.Println("Value of c is ", c, " ", "value of d is ", d)

	var e, f string = "hello", "world"
	fmt.Printf("Type of e, f are %T %T", e, f)
}
