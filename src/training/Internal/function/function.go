package main

import "fmt"

func MyFirstFunction() {
	fmt.Println("From: MyFirst function...")
}

func MyFunctionWithArg(str string) {
	fmt.Println("Input argument is ", str)
}

func MyFunc(str string, in int, fl float32) {
	fmt.Println("str:", str)
	fmt.Println("in:", in)
}

func Add(a int, b int) int {
	return a + b
}

func CircleProp(radius float64) (float64, float64) {

	area := 3.142 * radius * radius

	circ := 2 * 3.142 * radius

	return area, circ
}

func RectProps(width, height int) (area int, circ int) {
	area = width * height

	circ = 2 * (width + height)

	return
}

func main() {

	fmt.Println("Demo: Function")

	MyFirstFunction()

	MyFunctionWithArg("Hello")
	MyFunctionWithArg("I am good..")

	MyFunc("hello", 123, 33.44)

	fmt.Println("Add returned :", Add(10, 20))

	_, circ := CircleProp(10.23)
	fmt.Println(circ)
	area, cir := RectProps(10, 12)
	fmt.Println(area, cir)
}
