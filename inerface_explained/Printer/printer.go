package printer

import "fmt"

type Printer interface {
	Print() bool
	PageSetup() bool
	Priview() bool
}

func PrintContent(a Printer) bool {

	//alkajdlskajd
	//alksjdsakld
	b := a.PageSetup()
	c := a.Print()

	fmt.Println(a, b)

	return c || b
}
