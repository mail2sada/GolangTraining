package main

import "fmt"

type Employee struct {
	eid      int
	name     string
	dept     string
	location string
}

type ITResource struct {
	resourceID   int
	resourceName string
	Employee
}

// func (it ITResource) Print() {

// 	fmt.Println("=====================")
// 	fmt.Println("ResourceID ", it.resourceID)
// 	fmt.Println("Resource Name", it.resourceName)
// 	fmt.Println("assigned")
// 	it.assigned.Print()
// }

func (emp Employee) Print() {
	fmt.Println("==================================")
	fmt.Println("Eid:", emp.eid)
	fmt.Println("name:", emp.name)
	fmt.Println("dept:", emp.dept)
	fmt.Println("location:", emp.location)
}

func (emp *Employee) Initialise(eid int, name, dept, location string) {
	emp.eid = eid
	emp.name = name
	emp.dept = dept
	emp.location = location
}

func main() {

	fmt.Println("Demo: Methods")
	var emp Employee

	emp.Initialise(23232, "Anand Kumar", "Radio", "Bangalore")

	var itr ITResource = ITResource{resourceID: 1212, resourceName: "Mac book", Employee: emp}

	itr.Print()

	emp.Print()

}
