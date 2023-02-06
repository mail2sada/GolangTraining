package main

import "fmt"

func main() {

	fmt.Println("Demonstration double pointers...")

	var a int = 10

	var ptr *int = &a
	var dblPtr **int = &ptr

	fmt.Println("Value of A is ", a)
	fmt.Println("Value of A using pointer dererencing ", *ptr)
	fmt.Println("Value of A using double pointer dereferencing ", **dblPtr)

	fmt.Println("Address of A using & ", &a)
	fmt.Println("Address of A using ptr ", ptr)
	fmt.Println("Address of A using dblPtr ", *dblPtr)

	fmt.Println("Address of ptr using ptr ", &ptr)
	fmt.Println("Address of ptr using double pointer ", dblPtr)

	fmt.Println("Updating values using pointer")

	*ptr = 100

	fmt.Println("[AFTER_UPDATE PTR]Value of A is ", a)
	fmt.Println("[AFTER_UPDATE PTR]Value of A using pointer dererencing ", *ptr)
	fmt.Println("[AFTER_UPDATE PTR]Value of A using double pointer dereferencing ", **dblPtr)

	**dblPtr = 500

	fmt.Println("[AFTER_UPDATE DBL PTR]Value of A is ", a)
	fmt.Println("[AFTER_UPDATE DBL PTR]Value of A using pointer dererencing ", *ptr)
	fmt.Println("[AFTER_UPDATE DBL PTR]Value of A using double pointer dereferencing ", **dblPtr)

}
