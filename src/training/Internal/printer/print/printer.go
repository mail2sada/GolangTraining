package printer

type Printer interface {
	PrintPriview() []byte
	PageSetup() []byte
	Print() []byte
}

func PageSetup(art Printer) bool {

	art.PageSetup()

	return true
}

func PrintPriview(art Printer) bool {
	art.PrintPriview()
	return true
}

func Print(art Printer) bool {
	art.Print()
	return true
}
