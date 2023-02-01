package main

import "fmt"

func main() {
	defer demo_defer()
	fmt.Println("Demonstrait defer...")
	defer demo_defer_string("defer 1")

	defer demo_defer_string("defer 2")

	defer demo_defer_string("defer 3")

	demo_defer_string("this is from main...")
}

func demo_defer() {

	fmt.Println("Executing demo_defer...")
}

func demo_defer_string(str string) {
	fmt.Println(str)
}
