package main

import "fmt"

func iterateSlice(slice []int, fn func(value int)) {

	for i := 0; i < len(slice); i++ {

		fn(slice[i])

	}

}

func main() {

	fmt.Println("Demo: Functions....")

	func(a, b int) {
		fmt.Println("This is annonymous function...", a, b)
	}(10, 20)

	// fn := func(value int) {
	// 	fmt.Println("From annonymous function...")
	// 	fmt.Println(value)
	// }

	var sl []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	iterateSlice(sl, func(value int) {
		fmt.Println("From annonymous function...")
		fmt.Println(value)
	})
}
