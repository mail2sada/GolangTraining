package main

import "fmt"

func main() {

	a := 10

	defer fmt.Println("This is from defer...", a)

	TestMyDefer()

	defer fmt.Println("This is from defer 2")

	defer fmt.Println("This is from defer 3")

	a = 30
	fmt.Println("Demo: defer", a)
	fmt.Println("Hello...")
}

func TestMyDefer() {

	defer fmt.Println("This is from defer of TestMydefer...")
}



