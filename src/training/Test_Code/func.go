package main

import "fmt"

type test struct {
	name string
}

func main() {

	fmt.Println("Demo: Functions...")

	fmt.Println("Result of 10, 20 addition ", add(10, 20))

	fmt.Println("Result of 10, 20 addition ", sub(20, 10))

	fmt.Println(mul(10, 20))

	MyPrintln("121", "212142", "sdsdsds", "ddrdfd")

	MyCustomPrint(1, 2, "fadfa", 20.22, test{name: "test name"})

	var a, b interface{}

	a = 10

	b = "Hello"

	b = test{}

	fmt.Println(a, b)

	var myStringArray [10]string
	myArrayPrint(myStringArray)

	mySlicePrint(myStringArray[0 : len(myStringArray)-1])

}

func myStructPrint(t test) {
	
}

func myArrayPrint(arr [10]string) {

}

func mySlicePrint(slice []string) {

}

func add(a int, b int) int {

	c := a + b
	return c
}

func sub(a, b int) int {

	c := a + b
	return c
}

func mul(a, b int) (res int) {
	res = a * b
	return
}

func MyPrintln(elements ...string) {

	for _, val := range elements {
		fmt.Print(val)
	}
	fmt.Println()

}

func MyCustomPrint(elements ...interface{}) {

	for _, val := range elements {
		fmt.Print(val)
	}

}
