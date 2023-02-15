package main

import (
	"fmt"
	"reflect"
)

type author struct {
	name  string
	genre string
}

type book struct {
	bookid string
	title  string
	author
}

func main() {

	fmt.Println("Demo: Nested structure..")

	var bk book = book{bookid: "bk:1", title: "Golang", author: author{name: "Vishal", genre: "Computer Science"}}

	at := bk.author

	if at == bk.author {

	}
	if reflect.DeepEqual(at, bk.author) {

	}

	fmt.Println(bk)

	fmt.Println(bk.author)
}
