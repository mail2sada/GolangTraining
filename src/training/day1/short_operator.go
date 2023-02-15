package main

import (
	"fmt"
	"math"
)

func main() {


	

	a := 0.0

	fmt.Printf("Type of a is %T\n", a)

	b := 0

	fmt.Printf("Ty[e of b is %T", b)

	c := "hello world"

	fmt.Printf("Type of c is %T", c)

	d, e, f := 0.0, 0, "Test your skills"

	fmt.Printf("Type of d is %T\n Type of e is %T\n Type of f is %T", d, e, f)

	g := math.Sin(90)

	fmt.Printf("Type of g is %T", g)

	fmt.Println("Value of g is ", g)
}
