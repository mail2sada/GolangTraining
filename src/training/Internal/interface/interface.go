package main

import "fmt"

type interface_Name interface {
	v1()
	v2()
}

type student struct {
	roll_no int
	name    string
}

func (std student) v1() {

	fmt.Println("From student:v1()")

}

func (std student) v2() {
	fmt.Println("From student:v2()")
}

type emp struct {
	emp_id int
	name   string
}

func (std emp) v1() {
	fmt.Println("From emp:v1()")

}

func (std emp) v2() {
	fmt.Println("From emp:v2()")
}

func main() {

	fmt.Println("Demo: Interface")

	var a, b interface_Name

	a = student{}

	b = emp{}

	a.v1()

	b.v1()

	a = b

	a.v1()

}
