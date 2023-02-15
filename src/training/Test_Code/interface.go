package main

import "fmt"

type Printer interface {
	Print()

	PageSetup()

	Priview()
}

type Sturent struct {
	roll_no int
	name    string
	class   string
}

func (a Sturent) Print() {
	fmt.Println("In Print Sturent")

	fmt.Println(a)
}

func (a Sturent) PageSetup() {
	fmt.Println("In PageSetup Sturent")

}

func (a Sturent) Priview() {
	fmt.Println("In Priview Sturent")

}

type Teacher struct {
	tid    int
	name   string
	subjec string
}

func (a Teacher) Print() {

	fmt.Println("In Print Teacher")
	fmt.Println(a)
}

func (a Teacher) PageSetup() {
	fmt.Println("In PageSetup Teacher")
}

func (a Teacher) Priview() {
	fmt.Println("In Priview Teacher")
}

func main() {

	fmt.Println("Demo: Interface...")

	var std, tch Printer

	std = Sturent{roll_no: 1, name: "abcdef", class: "10"}

	tch = Teacher{tid: 1, name: "12345", subjec: "Maths"}

	// tch.PageSetup()
	// tch.Print()
	// tch.Priview()

	// std.Print()
	// std.Priview()
	// std.PageSetup()

	// tch = std

	// tch.PageSetup()
	// tch.Print()
	// tch.Priview()

	// a := std.(Sturent)

	// b := tch.(Sturent)

	// fmt.Println(a, b)

	GenericFuncForStudentTeacher(std)
	GenericFuncForStudentTeacher(tch)

}

func GenericFuncForStudentTeacher(a Printer) {

	a.Print()

	switch a.(type) {

	case Sturent:
		fmt.Println("Student...")

	case Teacher:
		fmt.Println("Teacher...")
	}

}
