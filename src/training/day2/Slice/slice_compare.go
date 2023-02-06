package main

import (
	"bytes"
	"fmt"
	"reflect"
)

//explore byte.Equal()

func main() {

	fmt.Println("Demonstraiting ")

	slice_1 := []byte{'M', 'A', 'V', 'E', 'N', 'I', 'R'}
	slice_2 := []byte{'M', 'A', 'V', 'E', 'N', 'I', 'R'}
	slice_3 := []byte{'M', 'A', 'V', 'E', 'n', 'I', 'R'}

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceInt := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(reflect.DeepEqual(slice, sliceInt))

	fmt.Println(reflect.DeepEqual(slice_1, slice_2))

	// Comparing slice
	// Using Compare function
	res := bytes.Compare(slice_1, slice_2)

	/*
		If the result is 0, then slice_1 == slice_2.
		If the result is -1, then slice_1 < slice_2.
		If the result is +1, then slice_1 > slice_2.
	*/

	if res == 0 {
		fmt.Println("!..Slices 1 & 2are equal..!")
	} else {
		fmt.Println("!..Slice are not equal..!")
	}

	res = bytes.Compare(slice_1, slice_3)

	if res == 0 {
		fmt.Println("!.. Slices are 1 & 3 are equal")
	} else {
		fmt.Println("!.. Slices are 1 & 3 are not equal")
	}

	fmt.Println("Value of res is ", res)
}
