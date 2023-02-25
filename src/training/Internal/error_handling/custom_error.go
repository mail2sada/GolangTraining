package main

import "fmt"

type MyCustomError struct {
	ErrorNo      int
	ErrorMessage string
}

func (err MyCustomError) Error() string {
	errMsg := fmt.Sprintln("\nErrorNo:", err.ErrorNo, "\nErrorMessage:", err.ErrorMessage)
	return errMsg
}

func IntegerAt(elements []int, pos int) (val int, err error) {

	if pos >= len(elements) {
		err = &MyCustomError{ErrorNo: 1234, ErrorMessage: "Array index out of bound,,,,"}

		return
	}

	if pos < 0 {
		err = &MyCustomError{ErrorNo: 1256, ErrorMessage: "Invalid position and can't be negative"}
		return
	}

	val = elements[pos]

	err = nil
	return
}

func main() {

	fmt.Println("Demo Custom error..")
	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	val, err := IntegerAt(slice, -1)

	fmt.Println("Value is ", val, " and error is ", err)
}
