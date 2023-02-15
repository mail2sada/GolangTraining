package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Demo: Sort package")

	var testSlice []int = []int{5, 3, 8, 1, 9, 2, 0, 2}

	fmt.Println("Before Sort")

	for _, val := range testSlice {
		fmt.Println(val)
	}

	fmt.Println(testSlice)

	sort.Ints(testSlice)

	fmt.Println("After sort.....")

	for _, val := range testSlice {
		fmt.Println(val)
	}

	fmt.Println(testSlice)

	var testStringSlice []string = make([]string, 10, 15)

	testStringSlice = []string{"1312321", "asdasdas", "werewrer", "sfsfsdf"}

	fmt.Println(testStringSlice)

	sort.Strings(testStringSlice)

	fmt.Println(testStringSlice)

	pos := sort.SearchStrings(testStringSlice, "sfsfsdf")

	fmt.Println(pos, testStringSlice)

}
