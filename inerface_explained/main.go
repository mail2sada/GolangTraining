package main

import (
	"fmt"

	document "github.com/mail2sada/interface_explained/Document"

	spreadsheet "github.com/mail2sada/interface_explained/SpreadSheet"

	printer "github.com/mail2sada/interface_explained/Printer"
)

func main() {

	var d, s printer.Printer

	d = document.Document{}

	s = spreadsheet.SpeadSheet{}

	fmt.Println("11111111")

	printer.PrintContent(d)

	fmt.Println("11111111")

	printer.PrintContent(s)

	fmt.Println("11111111")

}
