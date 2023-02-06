package main

import "fmt"

type Author struct {
	author_name string
	books       []Book
}

type Book struct {
	book_id     int
	book_title  string
	publication string
	year        int
}

func (aut Author) print() {
	fmt.Println("===============================")
	fmt.Println("Author Name    :", aut.author_name)
	for idx, book := range aut.books {
		fmt.Println("----------------------------")
		fmt.Println("Sl No        :", idx)
		fmt.Println("----------------------------")
		fmt.Println("Book _id     :", book.book_id)
		fmt.Println("Book title   :", book.book_title)
		fmt.Println("Publication  :", book.publication)
		fmt.Println("Year		  :", book.year)
		fmt.Println("----------------------------")
	}
	fmt.Println("===============================")
}

func (auth *Author) update() {
	auth.author_name = "Kumar Anand"
}

func main() {
	fmt.Println("Demonstraiting struct methods...")

	auth_details := Author{author_name: "Anand Kumar",
		books: []Book{
			Book{book_id: 1,
				book_title:  "Go lang in 15 days",
				publication: "Bharat",
				year:        2023},
			Book{book_id: 2,
				book_title:  "Advanced Golang...",
				publication: "Mavenir...",
				year:        2022},
		},
	}

	auth_details.print()

	auth_details.update()

	auth_details.print()

}
