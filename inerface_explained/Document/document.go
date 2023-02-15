package document

import "fmt"

type Document struct {
	pages     int
	name      string
	author    string
	watermark string
}

func (a Document) Print() bool {

	fmt.Println("IN Print Method.. Document")
	return true

}

func (a Document) Priview() bool {
	fmt.Println("IN Priview Method.. Document")
	return true

}

func (a Document) PageSetup() bool {

	fmt.Println("IN PageSetup Method.. Document")
	return true

}
