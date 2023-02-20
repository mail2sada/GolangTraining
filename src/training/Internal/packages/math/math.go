package math

import "fmt"

var a int = 20

var B int = 30

var MySum = 100

var yourSum = 300

func init() {

	fmt.Println("I am in my init method")
}

func init() {

	fmt.Println("I am in my init2 method")
}

func init() {

	fmt.Println("I am in my init3 method")
}

func testPrin() {

	fmt.Println("Value of A is ", a)
}

func TestPrint() {
	fmt.Println(yourSum)

	testPrin()
}
