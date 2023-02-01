package main

import "fmt"

func main() {
	fmt.Println("Mu;ti dimentional array")

	array := [3][3]string{{"This ", "is ", "Golang"}, {"Training ", "conducted ", "in "}, {"Mavenir ", "System ", "Private ltd"}}

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(array[x][y])
		}
	}

	var arrayInt [2][2]int
	arrayInt[0][0] = 1
	arrayInt[0][1] = 2
	arrayInt[1][0] = 3
	arrayInt[1][1] = 4

	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			fmt.Print(arrayInt[x][y])
		}
		fmt.Println("\n")
	}

}
