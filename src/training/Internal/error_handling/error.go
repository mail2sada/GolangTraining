package main

import (
	"errors"
	"fmt"
)

func IntegerAt(intSlice []int, pos int) (val int, err error) {

	if pos >= len(intSlice) {
		err = errors.New("Slice index out of bound exceptions")

		err = fmt.Errorf("This my test error %v", err)
		return
	}

	val = intSlice[pos]
	err = nil

	return
}

func main() {

	fmt.Println("Demo: Errors")

	var slice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	val, err := IntegerAt(slice, 100)

	if err != nil {

		fmt.Println("Received error while executing IntegerAt ", err)
	}

	fmt.Println("Value is ", val, " error is ", err)
}
