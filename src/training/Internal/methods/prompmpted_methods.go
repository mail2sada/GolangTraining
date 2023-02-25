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
	Address
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

func (e *Emp) setValue(eid int, name string, addr Address) {
	e.eid = eid
	e.name = name
	e.Address = addr
}

func main() {

	addr := Address{}

	emp := Emp{}

	addr.setValue("lashdklashdjl", "asldlkjd", 560099)

	emp.setValue(1234, "hello", addr)

	a, b, c := emp.getValue()

	fmt.Println(a, b, c)

}
