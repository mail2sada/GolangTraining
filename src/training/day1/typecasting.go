package main

import "fmt"

func main() {

	// taking the required
	// data into variables
	var totalsum int = 846
	var number int = 19
	var avg float32

	// explicit type conversion
	avg = float32(totalsum) / float32(number)

	// Displaying the result
	fmt.Printf("Average = %f\n", avg)
}
