package main

import "fmt"

func main() {

	fmt.Println("Demonstraing... annonynous functions")

	func() {
		fmt.Println("This is annonymous function..")
	}()

	annonym_func := func(a, b int) int {
		return a + b
	}

	a := func(a, b int) int {
		return (a + b)
	}(10, 20)

	fmt.Println("Value of a is ", a)

	fmt.Println(annonym_func(10, 20))
}
