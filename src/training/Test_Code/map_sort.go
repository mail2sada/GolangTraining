package main

import (
	"fmt"
	"sort"
)

// store all map keys in slice
// sort keys slice

//iterate through key and print mao values

func main() {

	fmt.Println("Demo: Sort keys in map...")

	var myMap = map[string]string{
		"Beta":  "A",
		"Alpha": "D",
		"Gama":  "C",
		"Delta": "E",
		"Penta": "X",
		"Hexa":  "Y",
	}

	var keySlice = []string{}

	for key, _ := range myMap {

		keySlice = append(keySlice, key)
	}

	sort.Strings(keySlice)

	for _, val := range keySlice {

		fmt.Println("key", val, "Value in key", myMap[val])
	}

}
