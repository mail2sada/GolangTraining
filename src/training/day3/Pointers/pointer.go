package main

import "fmt"

func main() {

	fmt.Println("Demonstration of pointers....")

	var a int = 10

	var p *int = &a

	fmt.Println("Value of A is ", a)

	fmt.Println("Address of A is ", &a)

	fmt.Println("Address of A with pointer is ", p)

	fmt.Println("Value of A with pointer dereferenceing ", *p)

	*p = 50

	fmt.Println("Value of A after updating using pointer ", a)

	fmt.Println("Value of A with pointer dereferencing ", *p)

	a = 100

	fmt.Println("Value of A after updating using pointer ", a)

	fmt.Println("Value of A with pointer dereferencing ", *p)

}
