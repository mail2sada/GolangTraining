package main

import (
	"fmt"
)

type person interface {
	update_name(name string)
	update_id(id int)
	print_details()
}

type employee struct {
	emp_id   int
	emp_name string
}

type student struct {
	roll_no int
	name    string
}

func (a *employee) update_name(name string) {
	a.emp_name = name

}
func (a *employee) update_id(id int) {
	a.emp_id = id

}
func (a *employee) print_details() {
	fmt.Println("Details of employee...")
	fmt.Println(a)
}

func (a *student) update_name(name string) {
	a.name = name
}
func (a *student) update_id(id int) {
	a.roll_no = id
}
func (a *student) print_details() {
	fmt.Println("Details of student...")
	fmt.Println(a)
}

func main() {

	fmt.Println("Demo: Interfaces....")
	var a, b person

	a = new(employee)
	b = new(student)

	var list []person = []person{a, b}

	a.update_id(10)
	a.update_name("Shivakumar T")

	b.update_id(20)

	a.update_name("Vishal p")

	for _, val := range list {

		val.print_details()

		switch val.(type) {

		case *student:
			fmt.Println("Val is of type student")
		case *employee:
			fmt.Println("Val is of type employee")
		}
	}

}
