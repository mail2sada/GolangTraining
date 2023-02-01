package main

import "fmt"

func main() {
	fmt.Println("Demonstraiting returning value name")

	a, b := area(5, 15)

	fmt.Println("Area of square and rectangle are ", a, b)

}

func area(p, q int) (square int, rectangle int) {

	square = p * p
	rectangle = p * q
	return
}
