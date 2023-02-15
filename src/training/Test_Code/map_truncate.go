package main

import "fmt"

func main() {

	fmt.Println("Demo: Truncate maps....")

	var myMap = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	fmt.Println(myMap)

	// for key, val := range myMap {

	// 	fmt.Println("key:", key, "val", val)

	// 	delete(myMap, key)
	// }

	myMap = make(map[string]int)

	fmt.Println(myMap, "============")
}
