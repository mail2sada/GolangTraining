package main

import "fmt"

func main() {

	fmt.Println("Demonstrating Slice...")
	var intSlice []int = []int{1, 2, 3, 4, 5}

	fmt.Println(intSlice)

	var intSlice1 = make([]int, 3)

	len := copy(intSlice1, intSlice)

	fmt.Println("Copied len...", len)

	fmt.Println(intSlice1, intSlice)

}
