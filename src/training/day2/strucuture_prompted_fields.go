package main

import "fmt"

type details struct {
	slno        int
	description string
}

type item struct {
	inventry_id int
	details
}

func main() {

	var items = item{inventry_id: 1, details: details{1, "HP printer..."}}

	fmt.Println(items)

}
