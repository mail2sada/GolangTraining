package main

import "fmt"

func main() {
	fmt.Println("Demonstraiting array argument to function...")
	var arr [6]int = [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	//printArray(arr1)
	printArray(arr)

	fmt.Printf("\n")

}

func printArray(arr [6]int) {

	for idx, val := range arr {
		fmt.Printf("\narr[%d] =  %d", idx, val)
	}
}
