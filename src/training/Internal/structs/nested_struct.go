package main

import "fmt"

type Address struct {
	Building string
	Street   string
	City     string
	pincode  int
}

type Employee struct {
	Name string
	Eid  string
	addr Address
}

func main() {

	fmt.Println("Demo nested struct")

	var emp Employee = Employee{Name: "Suresk BK", Eid: "208818", addr: Address{Building: "Bharatiya City", Street: "Hennur Road", City: "Bangalore", pincode: 560045}}

	addr := emp.addr

	fmt.Println(emp)

	fmt.Println(addr)

	fmt.Println(emp.addr.City)

	emp.addr.Building = "Sobha Daffodils"

	fmt.Println(emp)
}
