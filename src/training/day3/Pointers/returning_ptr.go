package main

import "fmt"

func main() {

	fmt.Println("Demonstrating returning pointer from function")

	a := function()

	fmt.Printf("Type of a is %T\n", a)

	fmt.Printf("Value of a is %d\n ", *a)

}

func function() *int {
	a := 10 * 10

	return &a
}
