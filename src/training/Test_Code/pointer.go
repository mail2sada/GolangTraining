package main

import "fmt"

func main() {

	fmt.Println("Demo: Pointers")

	a := 10

	var ptr *int

	ptr = &a

	fmt.Println("Value of a is ", a)
	fmt.Println("Address of a is ", &a)
	fmt.Println("Address of a using ptr ", ptr)
	fmt.Println("Value of a using ptr ", *ptr)

	var ptrPtr **int

	ptrPtr = &ptr

	fmt.Println(ptrPtr)
	fmt.Println(*ptrPtr)
	fmt.Println(**ptrPtr)

	var str string = "Welcome to Mavenir Golang training..."

	strPtr := &str

	fmt.Println("Address of str using &", &str)

	fmt.Println("Address of str using strPtr", strPtr)

	fmt.Println("Value of str using ptr us ", *strPtr)

}
