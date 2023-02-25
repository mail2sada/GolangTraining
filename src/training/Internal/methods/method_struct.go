package main

import "fmt"

type Address struct {
	name   string
	street string
	pin    int
}

type Emp struct {
	eid  int
	name string
}

func (a *Address) setValue(name string, street string, pin int) {
	a.name = name
	a.street = street
	a.pin = pin
}

func (a *Address) getValue() (name string, street string, pin int) {

	name = a.name
	street = a.street
	pin = a.pin
	return
}

func (e *Emp) setValue(eid int, name string) {
	e.eid = eid
	e.name = name
}

func main() {

	fmt.Println("Demo: methods structure...")

	var addr Address = Address{}

	addr.setValue("Vishal", "Bangalore", 560001)

	name, street, pin := addr.getValue()

	fmt.Println("Values of name:", name)
	fmt.Println("Values of street:", street)
	fmt.Println("Values of pin:", pin)

	var emp Emp = Emp{}

	emp.setValue(1234, "Anil Kumar")

	fmt.Println(emp)
}
