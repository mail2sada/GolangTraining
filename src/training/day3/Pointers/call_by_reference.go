package main

import "fmt"

func main() {
	fmt.Println("demonstraiting call by reference")

	var a, b int = 10, 20

	fmt.Println("Before calling swap ", a, b)

	swap(&a, &b)
	fmt.Println("Before calling swap ", a, b)
}

func swap(a, b *int) {

	c := *a
	*a = *b
	*b = c
}
