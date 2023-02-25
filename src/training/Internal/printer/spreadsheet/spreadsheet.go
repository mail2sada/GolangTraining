package spreadsheet

import "fmt"

type SpreadSheet struct {
	Name       string
	Sheet_id   int
	Author     string
	Sheets     int
	CellWidth  int
	CellHeight int
}

func (d SpreadSheet) PrintPriview() []byte {
	fmt.Println("In PrintPriview SpreadSheet")
	var data []byte
	return data

}
func (d SpreadSheet) PageSetup() []byte {
	fmt.Println("In PageSetup SpreadSheet")
	var data []byte
	return data

}
func (d SpreadSheet) Print() []byte {
	fmt.Println("In Print SpreadSheet")
	var data []byte
	return data
}
