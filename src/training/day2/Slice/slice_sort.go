package main

import (
	"fmt"
	"sort"
)

// sort, search, reverse, IntsAreSorted
func main() {

	fmt.Println(" Demonstraiting Slice sort")

	sliceInt_1 := []int{1, 2, 3, 5, 6, 4, 7, 8, 9}

	sliceInt_2 := []string{"Mavenir", "Golang", "Training", "Bangalore"}

	fmt.Println(sliceInt_1)
	fmt.Println(sliceInt_2)
	sort.Ints(sliceInt_1)
	sort.Strings(sliceInt_2)
	fmt.Println("sssd", sliceInt_1)
	fmt.Println(sliceInt_2)

	sort.Sort(sort.Reverse(sort.IntSlice(sliceInt_1)))
	fmt.Println(sliceInt_1)

	sort.Sort(sort.Reverse(sort.StringSlice(sliceInt_2)))

	fmt.Println(sliceInt_2)

}
