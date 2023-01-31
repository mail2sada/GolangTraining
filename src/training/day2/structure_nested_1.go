package main

import "fmt"

type address struct {
	flat_no string
	street  string
	city    string
	pincode string
}

type authors struct {
	name           string
	books_authored int
	addr           address
}

type book struct {
	title       string
	publication string
	author      authors
}

func main() {

	fmt.Println("Demonstration of nested structure....")

	var book = book{
		title:       "Lets learn Golang...",
		publication: "Mavenir University",
		author: authors{
			name:           "Sameer Khanna",
			books_authored: 30,
			addr:           address{flat_no: "304", street: "HSR", city: "Bangalore", pincode: "560102"}},
	}

	fmt.Println(book)
	fmt.Println(book.author)
	fmt.Println(book.author.addr)

}
