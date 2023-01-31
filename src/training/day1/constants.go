package main

import "fmt"

const A = "Hello"

func main() {

	const B = 0.0

	fmt.Println(A, B)

	const C, D int = 20, 30

	const E, F = 0.0, "Go lang is cool"

	fmt.Println(C, "  ", D)

	fmt.Println(E, "  ", F)

}
