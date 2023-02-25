package main

import (
	"fmt"

	document "github.com/mail2sada/printer/documet"
	printer "github.com/mail2sada/printer/print"
	"github.com/mail2sada/printer/spreadsheet"
)

func main() {

	fmt.Println("Demo: Interfacess....")

	var P1, P2 printer.Printer

	P1 = document.Document{Name: "Test.doc", Author: "Suresh BK", Document_id: 1222, Pages: 12}

	P2 = spreadsheet.SpreadSheet{Name: "test.xls", Author: "Vishal", Sheet_id: 3244, Sheets: 10}

	fmt.Println("==========Priniting Document P1")
	printer.PageSetup(P1)
	printer.PrintPriview(P1)
	printer.Print(P1)

	fmt.Println("=========Printing Document P2")
	printer.PageSetup(P2)
	printer.PrintPriview(P2)
	printer.Print(P2)

}
