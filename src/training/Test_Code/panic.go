package main

import (
	"fmt"
	"runtime/debug"
)

func main() {

	var a, b = 10, 0

	fmt.Println(Add(a, b))

	x, y := Divide(a, b)

	fmt.Println("Main Exiting...", x, y)

}

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (c int, err error) {
	defer func() {
		a := recover()

		fmt.Println(a)
		debug.PrintStack()
	}()

	c = 0
	err = nil
	if b == 0 {

		panic("asndlnsadls")
	}
	return a / b, err
}
