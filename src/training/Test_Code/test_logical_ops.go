package main

import "fmt"

var a = 100

func main() {

	a := 10

	b := 20.2

	fmt.Println(a, b)

	//a = int(b)

	b = float64(a)

	fmt.Println(b)
}
