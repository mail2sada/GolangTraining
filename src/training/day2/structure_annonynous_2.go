package main

import "fmt"

type annonymous_feilds struct {
	int
	string
	float32
	uint
}

func main() {

	fmt.Println("Demonstraining annonymous feilds")
	annon_struct := annonymous_feilds{int: 10, string: "hello", float32: 10.33, uint: 10}

	fmt.Println(annon_struct)

	var annonymous annonymous_feilds = annonymous_feilds{10, "hello", 10.33, 10}

	fmt.Println(annonymous)

	if annon_struct == annonymous {

		fmt.Println("equal...")
	} else {
		fmt.Println("Not equal..")
	}

}
