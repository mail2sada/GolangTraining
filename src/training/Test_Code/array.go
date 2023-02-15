package main

import "fmt"

func main() {

	fmt.Println("Demo: Arrays ")

	var intArray [5]int

	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3

	fmt.Println(intArray[0])
	fmt.Println(intArray[1])
	fmt.Println(intArray[2])

	fmt.Println(intArray[3])
	fmt.Println(intArray[4])

	var arrayInt [5]int = [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arrayInt); i++ {

		fmt.Println(arrayInt[i])
	}

	for idx, _ := range arrayInt {

		fmt.Println("array at idx is", idx)

	}

	var arrayString [5]string = [5]string{"Welcome", "to", "Golang", "Training"}

	for i := 0; i < len(arrayString); i++ {
		fmt.Println(arrayString[i])
	}

	for idx, val := range arrayString {

		fmt.Println("array at idx is", idx, val)

	}

	fmt.Println("==============================================")

	var integerArray [5]int = [5]int{1: 10, 3: 30}

	for idx, val := range integerArray {

		fmt.Println("array at idx is", idx, val)

	}

	myArray := [...]string{"aasasa", "sdsfdfdf", "wwewewew", "fgdfggfd", 6: "Welcome"}

	for idx, val := range myArray {

		fmt.Println("array at idx is", idx, val)

	}

	testArray := myArray

	testArray[0] = "Helloworld...."

	for idx, val := range testArray {

		fmt.Println("array at idx is", idx, val)

	}

	for idx, val := range myArray {

		fmt.Println("array at idx is", idx, val)

	}

	fmt.Println("Copy array by reference....")

	test := &myArray

	test[0] = "Mavenir Systems...."

	for idx, val := range test {

		fmt.Println("array at idx is", idx, val)

	}

	for idx, val := range myArray {

		fmt.Println("array at idx is", idx, val)

	}

}
