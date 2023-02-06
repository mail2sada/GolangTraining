package main

import "fmt"

func main() {

	fmt.Println("Demonstration of Append in Slices")

	sliceInt := []int{3, 2, 1, 0, 4, 5}
	fmt.Println(sliceInt)

	sliceString := append([]string{"Hello", "India", "123"})

	sliceInt = append(sliceInt[0:3], sliceInt[4:]...)

	sliceInt = append(sliceInt, 6, 7, 8, 9, 10)

	fmt.Println(sliceInt)

	fmt.Println(sliceString)

}
