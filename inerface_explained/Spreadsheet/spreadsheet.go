package spreadsheet

import "fmt"

type SpeadSheet struct {
	pages     int
	name      string
	author    string
	watermark string
}

func (a SpeadSheet) Print()bool {
	fmt.Println("In Print method spread sheet")
	return true
}

func (a SpeadSheet) Priview()bool {
	fmt.Println("In Priview method spread sheet")
	return true

}

func (a SpeadSheet) PageSetup()bool {
	fmt.Println("In PageSetup method spread sheet")
	return true

}
