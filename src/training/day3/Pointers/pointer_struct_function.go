package main

import "fmt"

type Student struct {
	roll_no int
	name    string
	class   string
}

func printStudent(abc *Student) {

	fmt.Println("Student detials")
	fmt.Println("===========================================")
	fmt.Println("\tStudent Roll No :", abc.roll_no)
	fmt.Println("\tStudent Name    :", abc.name)
	fmt.Println("\tStudent Class   :", abc.class)
	fmt.Println("===========================================")
}

func updateStructCallByValue(std Student) {
	std.roll_no = 100
}

func updateStructCallByReference(std *Student) {
	std.roll_no = 100
}

func main() {

	fmt.Println("Demonstraiting Struct pointers to function")

	std := Student{roll_no: 1, name: "Anil Kumar", class: "10 B"}

	printStudent(&std)

	updateStructCallByValue(std)
	fmt.Println("After call by value...")
	printStudent(&std)
	updateStructCallByReference(&std)

	fmt.Println("After call by reference...")

	printStudent(&std)

}
