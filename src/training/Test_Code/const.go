package main

import "fmt"

func main() {

	const a string = "circle.."

	const b, c = 1, "test"

	fmt.Printf("%T   %T", b, c)
}
