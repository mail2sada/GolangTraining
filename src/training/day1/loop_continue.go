package main

import "fmt"

func main() {
	var x int = 0

	// for loop work as a while loop
	for x < 8 {
		if x == 5 {

			// skip two iterations
			x = x + 2
			continue
		}
		fmt.Printf("value is: %d\n", x)
		x++
	}
}
