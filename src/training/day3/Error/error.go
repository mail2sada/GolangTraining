package main

import (
	"errors"
	"fmt"
)

type Student struct {
	roll_no  int
	name     string
	class    string
	subjects string
}

func validate_arguments(a Student) (Student, error) {
	std := a

	var err error = nil
	if a.roll_no == 0 {
		err = errors.New("invalid roll number")
	} else if a.name == "" {
		err = errors.New("invalid name")
	} else if a.class == "" {
		err = errors.New("invalid class")
	} else if a.subjects == "" {
		err = errors.New("invalid subjects")
	}

	return std, err
}

func main() {
	a := Student{roll_no: 1, name: "Vishal Kumar", class: "10th B"}
	fmt.Println("Demo: Usage of error interface.....")
	b, err := validate_arguments(a)
	fmt.Println("returned struct is ", b, err)
}
