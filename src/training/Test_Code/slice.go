package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Demo: Slice")

	var mySlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < len(mySlice); i++ {

		fmt.Println(mySlice[i])
	}

	for idx, val := range mySlice {

		fmt.Printf("Value at mySlice[%d] = %d\n", idx, val)
	}

	mySlice = append(mySlice, 10)

	mySlice = append(mySlice)

	fmt.Println("After appending")

	fmt.Println("capacity of myslice is ", cap(mySlice))

	for idx, val := range mySlice {

		fmt.Printf("Value at mySlice[%d] = %d\n", idx, val)
	}

	idx := 4

	mySlice = append(mySlice[:idx], mySlice[idx+1:]...)

	fmt.Println(mySlice)

	var my2Slice []int

	sort.Ints(mySlice)

	copy(my2Slice, mySlice) //Copying slice

}
