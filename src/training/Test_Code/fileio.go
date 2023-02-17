package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Demo OS Package....")

	os.Rename("abc.csv", "data.csv")

}
/Users/sadanandd/Documents/GitHub/GolangTraining/src/training/Test_Code/fileio.go