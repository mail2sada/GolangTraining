package main

import "fmt"

type student struct {
	roll_no int
	name    string
	class   string
}

type teacher struct {
	emp_id  int
	name    string
	subject string
}

func main() {
	fmt.Println("... Demonstraiting methods with same name..")

	var std student

	std = std.setvalues(10, "Hemant Kulkarni..", "B Tech...")
	std.print()

	var tch teacher
	tch = tch.setvalues(20, "Vinay bhushan..", "Data Science..")

	tch.print()
}

func (std student) print() {
	fmt.Println(std)
}

func (std student) setvalues(a int, name string, class string) student {
	std.roll_no = a
	std.name = name
	std.class = class
	return std
}

func (tch teacher) print() {
	fmt.Println(tch)
}

func (tch teacher) setvalues(a int, name string, subject string) teacher {
	tch.emp_id = a
	tch.name = name
	tch.subject = subject
	return tch
}
