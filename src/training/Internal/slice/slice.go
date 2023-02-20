package main

import "fmt"

func main() {
	fmt.Println("Demo: Slice")

	var mySlice []int = []int{1, 2, 3, 4, 5}

	for idx, val := range mySlice {

		fmt.Println(idx, val)
	}

	for i := 0; i < len(mySlice); i++ {

		fmt.Println(mySlice[i])
	}
	fmt.Println("====================", cap(mySlice))

	mySlice = append(mySlice, 10, 20, 30, 40, 50, 60)

	fmt.Println("====================", cap(mySlice))

	for idx, val := range mySlice {

		fmt.Println(idx, val)
	}

	mySlice = append(mySlice[0:5], mySlice[6:]...)
	fmt.Println("====================", cap(mySlice))
	for idx, val := range mySlice {

		fmt.Println(idx, val)
	}
}
