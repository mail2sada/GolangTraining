package main

import "fmt"

func main() {

	fmt.Println("Demo: Defer")

	defer defer_demo()

	defer fmt.Println("Executing first defer...")

	defer fmt.Println("Executing second defer...")

}

func defer_demo() {
	fmt.Println("This is for demo..")
	defer fmt.Println("Calling this from defer_demo")
}
