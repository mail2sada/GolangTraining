package main

import "fmt"

func main() {
	fmt.Println("Slice as argument to function")

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)

	updateSlice(slice)

	fmt.Println(slice)

}

func updateSlice(sli []int) {
	sli[6] = 100
}
