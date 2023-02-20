package main

import "fmt"

func main() {

	fmt.Println("Demo: Arrays.")

	var myArray [10]int

	myArray[0] = 1

	myArray[1] = 2

	c := myArray[4]

	fmt.Println(c)

	for i := 0; i < len(myArray); i++ {

		fmt.Println("Value at index i is ", myArray[i])
	}

	var myTestArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var myArrayWt = [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for i := 0; i < len(myTestArray); i++ {

		fmt.Println("Value at index i is ", myTestArray[i])
	}

	for i := 0; i < len(myArrayWt); i++ {

		fmt.Println("Value at index i is ", myArrayWt[i])
	}

	for idx, val := range myArrayWt {
		fmt.Println(idx, val)
	}

	var strArray = [10]string{3: "Hello", 1: "my", 6: "dear", 4: "friends"}

	for idx, val := range strArray {
		fmt.Println(idx, val)
	}

	for i := 0; i < len(strArray); i++ {

		fmt.Println("Value at index i is ", strArray[i])
	}

}
