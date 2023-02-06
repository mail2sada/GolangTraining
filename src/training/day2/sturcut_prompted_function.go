package main

import "fmt"

type details struct {
	name   string
	age    int
	gender string
	salary int
}

type employee struct {
	eid         string
	designation string
	details
}

func (d details) promptedFunction() {
	fmt.Println("Employee details...")
	fmt.Println(d)
}

func main() {
	fmt.Println("Demonstraion of prompted functions")

	emp_details := employee{eid: "12345", designation: "Sr Manager", details: details{name: "Vasu M", age: 35, gender: "Male", salary: 3500000}}
	emp_details.promptedFunction()
}
