package main

import "fmt"

func main() {

	fmt.Println("Demo: Merge maps..")

	var a = map[string]int{"a": 2, "c": 3, "e": 4}

	var b = map[string]int{"a": 5, "b": 3, "d": 4}

	for key, val := range b {

		a[key] = val
	}

	fmt.Println(a)

	fmt.Println(b)

}
