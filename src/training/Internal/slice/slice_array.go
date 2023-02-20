package main

import "fmt"

func main() {

	fmt.Println("Demo: Slice with Array..")

	var intArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sliceFromArray := intArray[2:6]

	fmt.Println(sliceFromArray)

	fmt.Println(cap(sliceFromArray))

	sliceFromArray[7] = 100

	for i := 0; i <= len(intArray); i++ {
		fmt.Println(intArray[i])
	}

}
