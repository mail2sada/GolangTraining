package documet

import "fmt"

type Document struct {
	Name        string
	Document_id int
	Author      string
	Pages       int
	Width       int
	Height      int
}

func (d Document) PrintPriview() []byte {

	fmt.Println("In PrintPriview Document")

	var data []byte
	return data

}
func (d Document) PageSetup() []byte {
	fmt.Println("In PageSetup Document")
	var data []byte
	return data

}
func (d Document) Print() []byte {
	fmt.Println("In Print Document")
	var data []byte
	return data
}
