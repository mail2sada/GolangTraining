package main

import "fmt"

func main() {
	fmt.Println("Demo: loop")

	// for i := 0; i <= 10; i++ {

	// 	fmt.Println(i)
	// }

	// i := 1

	// for {
	// 	fmt.Println(i)

	// 	if i%10 == 0 {
	// 		i += 1
	// 		continue

	// 	}
	// 	i += 1
	// 	if i == 50 {
	// 		break
	// 	}

	// }

	var test string = "Mavenir Golang traiining...."

	for a, b := range test {

		fmt.Printf("%d == [%c]\n", a, b)
	}

}
