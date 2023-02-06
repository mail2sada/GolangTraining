package main

import "fmt"

// IMPORTANT FUNCTIONS IN SLICE

// len, cap, append, copy

func main() {
	fmt.Println("Slice demonstration")

	var intSlice []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var stringSlice []string = []string{"A", "B", "C", `D`, "E", "F"}

	fmt.Println(intSlice)

	fmt.Println(sliceInt)

	fmt.Println(stringSlice)
}
